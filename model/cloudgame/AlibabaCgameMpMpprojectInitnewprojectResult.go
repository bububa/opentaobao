package cloudgame

// AlibabaCgameMpMpprojectInitnewprojectResult 结构体
type AlibabaCgameMpMpprojectInitnewprojectResult struct {
	// 0
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// sucess
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// login session
	Data *MpProjectConfigDto `json:"data,omitempty" xml:"data,omitempty"`
}
