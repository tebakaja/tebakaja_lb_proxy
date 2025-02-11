package national_currency

import (
	"fmt"
	"context"
	"net/http"
	"encoding/json"

	helpers "tebakaja_lb_proxy/proxy/helpers"
)


/*
 *  --- National Currency Prediction Model Lists Service ---
 */
func (s *NationalCurrencyServiceImpl) NationalCurrencyListsService(ctx context.Context) (ApiResponse, error) {
	endpoint := fmt.Sprintf("%s/lists", helpers.GetEndpointService("national"))
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return ApiResponse{
			Message:    fmt.Sprintf("Failed to make request: %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ApiResponse{
			Message:    fmt.Sprintf("Failed to make request: %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ApiResponse{
			Message:    fmt.Sprintf("Request failed: %d", resp.StatusCode),
			StatusCode: resp.StatusCode,
		}, fmt.Errorf("request failed: %d", resp.StatusCode)
	}

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return ApiResponse{
			Message:    fmt.Sprintf("Failed to parse JSON response: %v", err),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return apiResponse, nil
}
