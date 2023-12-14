package bankindonesia

import (
	cfg "exchangerate/config"
	"exchangerate/dto"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

const source dto.ScaperSource = "bankindonesia"

type BankIndonesia struct {
}

func New() *BankIndonesia {
	return &BankIndonesia{}
}

func (s *BankIndonesia) Run(c *colly.Collector) (*dto.ScaperResponse, error) {
	sic := &dto.ScaperItemCollection{}
	ch := make(chan string)

	c.OnHTML("#ctl00_PlaceHolderMain_g_6c89d4ad_107f_437d_bd54_8fda17b556bf_ctl00_GridView1 table tbody tr", func(el *colly.HTMLElement) {
		data := el.ChildTexts("td")

		sic.AddItem(&dto.ScaperItem{
			BaseCurrencyCode:     dto.CURRENCY_CODE_IDR,
			Source:               source,
			TransferCurrencyCode: data[0],
			PriceBuy:             formatAmount(data[3]),
			PriceSell:            formatAmount(data[2]),
			Value:                formatAmount(data[1]),
		})
	})

	c.OnHTML("#tableData .search-box-wrapper", func(el *colly.HTMLElement) {
		date := el.ChildText("span")
		ch <- date
	})

	date, ok := <-ch
	if !ok {
		return &dto.ScaperResponse{}, fmt.Errorf("error: channel closed before receiving the value")
	}

	return sic.ToResponse(&dto.ScaperResponseArgs{
		TransferDate: formatDate(date),
	}), nil
}

func (s *BankIndonesia) Uri() string {
	return cfg.Spider.BankIndonesiaUri
}

func formatAmount(value string) float64 {
	cleanedValue := strings.ReplaceAll(value, ",", "")
	amount, err := strconv.ParseFloat(cleanedValue, 64)

	if err != nil {
		return 0
	}

	return amount
}

func formatDate(value string) string {
	date, err := time.Parse(dto.DATE_FORMAT_WEB, value)

	if err != nil {
		date = time.Now()
	}

	return date.Format(dto.DATE_FORMAT_DB)
}
