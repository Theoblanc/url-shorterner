package shortener

// ShortenRepositoryImplemnt get ShortenRepository
type ShortenRepositoryImplemnt struct {
	shortenRepository ShortenRepository
}

// Service is interface fo shortener
type Service interface {
	Create(dto *CreateShortenDTO) (string, error)
	GetAll() error
	GetByURL(url string) (GetShorteerDTO, error)
}

// Create create shortener service
func (sh *ShortenRepositoryImplemnt) Create(dto *CreateShortenDTO) (string, error) {
	tx := sh.shortenRepository.Save(dto)
	return tx, nil
}

// GetAll get all shortener (TODO::)
func (sh *ShortenRepositoryImplemnt) GetAll() {
	return
}

// GetByURL get shortener by url
func (sh *ShortenRepositoryImplemnt) GetByURL(url string) (GetShorteerDTO, error) {
	shoterner, err := sh.shortenRepository.FindByURL(url)

	return domainToDto(shoterner), err
}

func domainToDto(domain Anemic) GetShorteerDTO {
	return GetShorteerDTO{
		ID:  domain.ID,
		URL: domain.URL,
	}
}
