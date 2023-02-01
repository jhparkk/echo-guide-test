package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type RequestStat struct {
	Uptime       time.Time      `json:"uptime"`
	RequestCount uint64         `json:"requestCount"`
	Methods      map[string]int `json:"statuses"`
	mutex        sync.RWMutex
}

func NewRequestStat() *RequestStat {
	return &RequestStat{
		Uptime:  time.Now(),
		Methods: map[string]int{},
	}
}

func (s *RequestStat) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()

		s.RequestCount++
		method := c.Request().Method
		s.Methods[method]++

		return nil
	}
}

func (s *RequestStat) Stats(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}
