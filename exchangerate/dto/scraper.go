package dto

import (
	u "exchangerate/utils"
	"time"
)

const (
	CURRENCY_CODE_IDR string = "IDR"
	CURRENCY_CODE_JPY string = "JPY"
	CURRENCY_CODE_USD string = "USD"
	DATE_FORMAT_DB    string = time.DateOnly
	DATE_FORMAT_WEB   string = "2 January 2006"
)

type ScaperSource string

type ScaperArgs struct {
}

type ScaperItemCollection []*ScaperItem

type ScaperItem struct {
	BaseCurrencyCode     string
	PriceBuy             float64
	PriceSell            float64
	TransferCurrencyCode string
	TransferDate         string
	TransferRate         float64
	Value                float64
	Source               ScaperSource
}

type ScrapeRateArgs struct {
	Code        string
	DefaultRate float64
}

type ScaperResponseArgs struct {
	TransferDate string
}

type ScaperResponse []*Scaper

type Scaper struct {
	BaseCurrencyCode     string
	TransferCurrencyCode string
	TransferDate         string
	TransferRate         float64
	Source               ScaperSource
}

func (sic *ScaperItemCollection) AddItem(si *ScaperItem) {
	*sic = append(*sic, si)
}

func (sic *ScaperItemCollection) ToResponse(args *ScaperResponseArgs) *ScaperResponse {
	sr := ScaperResponse{}

	for _, item := range *sic {
		sr.addItem(&Scaper{
			BaseCurrencyCode:     item.BaseCurrencyCode,
			Source:               item.Source,
			TransferCurrencyCode: item.TransferCurrencyCode,
			TransferDate:         args.TransferDate,
			TransferRate: sic.findRate(&ScrapeRateArgs{
				Code:        item.TransferCurrencyCode,
				DefaultRate: item.Value,
			}),
		})
	}

	return &sr
}

func (sic *ScaperItemCollection) findItem(code string) (*ScaperItem, bool) {
	for _, si := range *sic {
		if si.TransferCurrencyCode == code {
			return si, true
		}
	}

	return &ScaperItem{}, false
}

func (sic *ScaperItemCollection) findRate(args *ScrapeRateArgs) float64 {
	si, exists := sic.findItem(args.Code)

	if !exists {
		return 1
	}

	return si.calculateRate(args.DefaultRate)
}

func (si *ScaperItem) calculateRate(precision float64) float64 {
	rate := (si.PriceBuy + si.PriceSell) / (2 * precision)

	if si.TransferRate != 0 {
		rate = si.TransferRate / precision
	}

	return u.Round(rate, 6)
}

func (sr *ScaperResponse) addItem(s *Scaper) {
	*sr = append(*sr, s)
}
