package shortener

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// ShortenerEntity shortener entity
type ShortenerEntity struct {
	ID          string `gorm:"primaryKey"`
	URL         string
	CustomShort string
	Expiry      time.Time
}

// ShortenerAnemic shortener anemic model
type ShortenerAnemic struct {
	ID          string
	URL         string
	CustomShort string
	Expiry      time.Time
}

// ShortenRepository is interface for shotrener
type ShortenRepository interface {
	Save(dto *CreateShortenDTO) string
	FindAll() (*ShortenerAnemic, error)
	FindByURL(url string) (ShortenerAnemic, error)
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
func (r *Repositorys) FindByURL(url string) (ShortenerAnemic, error) {
	shorten := ShortenerEntity{
		URL: url,
	}

	if cache := r.getCache(url); cache != nil {
		return convertEntityToDomain(*cache), nil
	}

	tx := r.DB.Find(&shorten, 1)

	if err := tx.Error; err != nil {
		return convertEntityToDomain(ShortenerEntity{}), err
	}
	r.setCache(url, &shorten)

	return convertEntityToDomain(shorten), nil
}

func convertDtoToEntity(dto CreateShortenDTO) ShortenerEntity {
	return ShortenerEntity{
		URL:         dto.url,
		CustomShort: dto.customShort,
		Expiry:      dto.expiry,
	}
}

func convertEntityToDomain(entity ShortenerEntity) ShortenerAnemic {
	return ShortenerAnemic{
		ID:          entity.ID,
		URL:         entity.URL,
		CustomShort: entity.CustomShort,
		Expiry:      entity.Expiry,
	}
}

func (r *Repositorys) getCache(
	key string,
) *ShortenerEntity {
	data, getDataFromRedisError := r.redis.Get("profile:" + key).Result()
	if getDataFromRedisError != nil {
		return nil
	}

	entity := &ShortenerEntity{}
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
	key string, shortenerEntity *ShortenerEntity,
) {
	marshaledEntity, _ := json.Marshal(&shortenerEntity)
	r.redis.Set(
		"profile:"+key, string(marshaledEntity), time.Second,
	)
}
