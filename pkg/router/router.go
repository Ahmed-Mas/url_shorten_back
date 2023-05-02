package router

import (
	"io"
	"log"
	"net/http"

	"github.com/Ahmed-Mas/url_shorten_back/pkg/serializer"
	"github.com/Ahmed-Mas/url_shorten_back/pkg/shorten"
	"github.com/gin-gonic/gin"
)

// todo: env var this
const DOMAIN = "http://localhost:10000/"

type Handlerer interface {
	RedirectToLong(*gin.Context)
	GenerateShort(*gin.Context)
}

type handler struct {
	service shorten.Shortener
}

func NewHandler(service shorten.Shortener) Handlerer {
	return &handler{service: service}
}

func (h *handler) serializer() shorten.Serializer {
	return &serializer.Serializer{}
}

func (h *handler) RedirectToLong(c *gin.Context) {
	log.Println("endpoint reached: redirect to long url")

	shortUrl := c.Params.ByName("url")
	urlData := h.service.Retrieve(shortUrl)

	log.Printf("redirecting to: %s\n", urlData.LongUrl)
	c.Redirect(http.StatusMovedPermanently, urlData.LongUrl)
}

func (h *handler) GenerateShort(c *gin.Context) {
	log.Println("endpoint reached: generate short url")

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err)
		return
	}

	s := h.serializer()
	urlData, err := s.Decode(body)
	if err != nil {
		c.Error(err)
		return
	}

	h.service.Store(urlData)

	urlData.ShortUrl = DOMAIN + urlData.ShortUrl

	respBody, err := s.Encode(urlData)
	if err != nil {
		c.Error(err)
		return
	}
	log.Println("returning shortened url")
	c.Data(200, "application/json", respBody)
}
