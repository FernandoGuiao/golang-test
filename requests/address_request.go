package requests

type CreateAddressRequest struct {
	Street     string `json:"street" binding:"required,min=3,max=100"`
	City       string `json:"city" binding:"required,min=2,max=50"`
	State      string `json:"state" binding:"required,len=2"`
	Country    string `json:"country" binding:"required,iso3166_1_alpha2"`
	Number     string `json:"number" binding:"required,max=20"`
	Complement string `json:"complement" binding:"max=50"`
}

type UpdateAddressRequest struct {
	Street     string `json:"street" binding:"omitempty,min=3,max=100"`
	City       string `json:"city" binding:"omitempty,min=2,max=50"`
	State      string `json:"state" binding:"omitempty,len=2"`
	Country    string `json:"country" binding:"omitempty,iso3166_1_alpha2"`
	Number     string `json:"number" binding:"omitempty,max=20"`
	Complement string `json:"complement" binding:"omitempty,max=50"`
}
