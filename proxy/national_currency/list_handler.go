package national_currency

import (
	"log"
	"sync"
	"time"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)



// NationalCurrencyListsHandler godoc
// @Summary            National Currency List
// @Description        National Currency List
// @Tags               National Currency
// @Produce            json
// @Success            200  {object} ApiResponse
// @Router             /national-currency/lists [get]
func NationalCurrencyListsHandler(service NationalCurrencyService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(c.Context(), 120*time.Second)
		defer cancel()

		ch := make(chan ApiResponse, 1)
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()

			apiResponse, err := service.NationalCurrencyListsService(ctx)
			if err != nil {
				log.Printf("[%s] %v", time.Now().Format("2006-01-02 15:04:05"), err)
				ch <- apiResponse
				return
			}

			ch <- apiResponse
		}()

		go func() {
			wg.Wait()
			close(ch)
		}()

		select {
			case apiResponse, ok := <-ch:
				if !ok {
					return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to get a response from the server",
					})
				}
				return c.Status(apiResponse.StatusCode).JSON(apiResponse)

			case <-ctx.Done():
				log.Printf("[%s] Timeout: %v", time.Now().Format("2006-01-02 15:04:05"), ctx.Err())
				return c.Status(http.StatusRequestTimeout).JSON(fiber.Map{
					"error": "Request timeout",
				})
		}
	}
}
