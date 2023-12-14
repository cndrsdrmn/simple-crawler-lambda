package spiders

import (
	cfg "exchangerate/config"
	"exchangerate/dto"
	bi "exchangerate/spiders/bankindonesia"
	"exchangerate/spiders/ortax"
	"log"

	"github.com/gocolly/colly/v2"
	ext "github.com/gocolly/colly/v2/extensions"
)

var spiders []Spider

type Spider interface {
	Run(c *colly.Collector) (*dto.ScaperResponse, error)
	Uri() string
}

func init() {
	spiders = []Spider{
		bi.New(),
		ortax.New(),
	}
}

func Exec() *dto.ScaperResponse {
	response := dto.ScaperResponse{}

	for _, s := range spiders {
		c := colly.NewCollector()

		// uncommend if using local file
		// if cfg.App.Debug {
		// 	t := &http.Transport{}
		// 	t.RegisterProtocol("file", http.NewFileTransport(http.Dir("/")))
		// 	c.WithTransport(t)
		// }

		if cfg.App.IsProduction() {
			ext.RandomUserAgent(c)
		}

		c.OnError(func(r *colly.Response, err error) {
			log.Println("Error", err.Error())
		})

		go func() {
			if err := c.Visit(s.Uri()); err != nil {
				log.Println("Error visiting URL:", err)
				return
			}
		}()

		res, err := s.Run(c)
		if err != nil {
			log.Println("Error scraping data from URL:", s.Uri())
		}

		response = append(response, *res...)
	}

	return &response
}
