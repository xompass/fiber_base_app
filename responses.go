package fiber_base_app

type Count struct {
	Count int64 `json:"count"`
} // @name CountResponse

type Exists struct {
	Exists bool `json:"exists"`
} // @name ExistsResponse

type CustomHTTPError struct {
	Code    int           `json:"code"`
	Message string        `json:"message,omitempty"`
	Details []interface{} `json:"details,omitempty"`
} // @name HTTPError

func (err CustomHTTPError) Error() string {
	return err.Message
}
