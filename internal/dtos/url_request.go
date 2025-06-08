package dtos

type URLRequest struct {
	URL             string `json:"url"`
	ExpireInMinutes int    `json:"expire_in_minutes"`
}
