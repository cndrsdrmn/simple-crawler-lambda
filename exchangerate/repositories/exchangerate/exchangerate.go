package exchangerate

import (
	"exchangerate/database"
	"exchangerate/dto"
	"exchangerate/model"
)

type ExchangeRateRepo struct {
}

func New() *ExchangeRateRepo {
	return &ExchangeRateRepo{}
}

func (repo *ExchangeRateRepo) Upsert(attrs *dto.ScaperResponse) error {
	newAttrs := []*model.ExchangeRate{}

	for _, attr := range *attrs {
		upd, err := repo.update(attr)
		if err != nil {
			return err
		}

		if !upd {
			newAttrs = append(newAttrs, &model.ExchangeRate{
				BaseCurrencyCode:     attr.BaseCurrencyCode,
				Source:               attr.Source,
				TransferCurrencyCode: attr.TransferCurrencyCode,
				TransferDate:         attr.TransferDate,
				TransferRate:         attr.TransferRate,
			})
		}
	}

	if len(newAttrs) > 0 {
		if err := repo.batchInsert(newAttrs); err != nil {
			return err
		}
	}

	return nil
}

func (repo *ExchangeRateRepo) batchInsert(attrs []*model.ExchangeRate) error {
	return database.Instance.
		Model(&model.ExchangeRate{}).
		Create(attrs).Error
}

func (repo *ExchangeRateRepo) update(attr *dto.Scaper) (bool, error) {
	var count int64
	query := database.Instance.
		Model(&model.ExchangeRate{}).
		Where(&model.ExchangeRate{
			BaseCurrencyCode:     attr.BaseCurrencyCode,
			Source:               attr.Source,
			TransferCurrencyCode: attr.TransferCurrencyCode,
			TransferDate:         attr.TransferDate,
		})

	query.Count(&count)

	if count > 0 {
		if err := query.Updates(&model.ExchangeRate{
			TransferRate: attr.TransferRate,
		}).Error; err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
