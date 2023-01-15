package shortener

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// Entity shortener entity
type Entity struct {
	ID          string `gorm:"primaryKey"`
	URL         string
	CustomShort string
	Expiry      time.Time
}

// Anemic shortener anemic model
type Anemic struct {
	ID          string
	URL         string
	CustomShort string
	Expiry      time.Time
}

// ShortenRepository is interface for shotrener
type ShortenRepository interface {
	Save(dto *CreateShortenDTO) string
	FindAll() (*Anemic, error)
	FindByURL(url string) (Anemic, error)
}

// Repositorys postgresql and redis repository
type Repositorys struct {
	*gorm.DB
	redis *redis.Client
}

// Save save given domain object
func (r *Repositorys) Save(dto *CreateShortenDTO) {
	r.DB.Save(convertDtoToEntity(*dto))
}

// FindByURL find shortener by url and save cache
func (r *Repositorys) FindByURL(url string) (Anemic, error) {
	shorten := Entity{
		URL: url,
	}

	if cache := r.getCache(url); cache != nil {
		return convertEntityToDomain(*cache), nil
	}

	tx := r.DB.Find(&shorten, 1)

	if err := tx.Error; err != nil {
		return convertEntityToDomain(Entity{}), err
	}
	r.setCache(url, &shorten)

	return convertEntityToDomain(shorten), nil
}

func convertDtoToEntity(dto CreateShortenDTO) Entity {
	return Entity{
		URL:         dto.url,
		CustomShort: dto.customShort,
		Expiry:      dto.expiry,
	}
}

func convertEntityToDomain(entity Entity) Anemic {
	return Anemic{
		ID:          entity.ID,
		URL:         entity.URL,
		CustomShort: entity.CustomShort,
		Expiry:      entity.Expiry,
	}
}

func (r *Repositorys) getCache(
	key string,
) *Entity {
	data, getDataFromRedisError := r.redis.Get("profile:" + key).Result()
	if getDataFromRedisError != nil {
		return nil
	}

	entity := &Entity{}
	jsonUnmarshalError := json.Unmarshal([]byte(data), entity)
	if jsonUnmarshalError != nil {
		return nil
	}

	if entity.ID == "" {
		return nil
	}
	return entity
}

func (r *Repositorys) setCache(
	key string, shortenerEntity *Entity,
) {
	r.redis.Set(
		"shotrener:"+key, string(shortenerEntity.CustomShort), time.Duration(shortenerEntity.Expiry.Hour()),
	)
}
