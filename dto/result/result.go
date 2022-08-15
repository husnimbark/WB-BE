package dto

type SuccessResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"users"`
}
type SuccessRegister struct {
	Code int         `json:"code"`
	Data interface{} `json:"user"`
}
type SuccessDeleteUser struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type SuccesFindProducts struct {
	Code int         `json:"code"`
	Data interface{} `json:"products"`
}
type SuccessGetProduct struct {
	Code int         `json:"code"`
	Data interface{} `json:"product"`
}
type SuccessDeleteProduct struct {
	Code int         `json:"code`
	Data interface{} `json:"data"`
}
type SuccesFindToppings struct {
	Code int         `json:"code"`
	Data interface{} `json:"toppings"`
}
type SuccesGetTopping struct {
	Code int         `json:"code"`
	Data interface{} `json:"topping"`
}
