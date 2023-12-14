package ortax

import (
	cfg "exchangerate/config"
	"exchangerate/dto"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const source dto.ScaperSource = "ortax"

type Ortax struct {
}

func New() *Ortax {
	return &Ortax{}
}

func (s *Ortax) Run(c *colly.Collector) (*dto.ScaperResponse, error) {
	sic := &dto.ScaperItemCollection{}
	ch := make(chan bool)

	c.OnHTML("#table-kursmk tbody tr", func(el *colly.HTMLElement) {
		data := el.ChildTexts("td")
		code := data[1]
		value := 1

		if code == dto.CURRENCY_CODE_JPY {
			value = 100
		}

		sic.AddItem(&dto.ScaperItem{
			BaseCurrencyCode:     dto.CURRENCY_CODE_IDR,
			Source:               source,
			TransferCurrencyCode: code,
			TransferRate:         formatAmount(el.ChildText("th")),
			Value:                float64(value),
		})
	})

	c.OnHTML("title", func(el *colly.HTMLElement) {
		ch <- true
	})

	_, ok := <-ch
	if !ok {
		return &dto.ScaperResponse{}, fmt.Errorf("error: channel closed before receiving the value")
	}

	return sic.ToResponse(&dto.ScaperResponseArgs{
		TransferDate: time.Now().Format(time.DateOnly),
	}), nil
}

func (s *Ortax) Uri() string {
	return cfg.Spider.OrtaxUri
}

func formatAmount(value string) float64 {
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, ",", ".")

	amount, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return amount
}
