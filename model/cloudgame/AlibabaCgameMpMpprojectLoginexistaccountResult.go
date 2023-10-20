package cloudgame

// AlibabacgamempmpprojectloginexistaccountResult 结构体
type AlibabacgamempmpprojectloginexistaccountResult struct {
	// 0
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// sucess
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// login session
	Data *LoginSessionDto `json:"data,omitempty" xml:"data,omitempty"`
}
