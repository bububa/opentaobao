package vaccin

// IsvVcAvailableRemindRequest 结构体
type IsvVcAvailableRemindRequest struct {
	// 阿里健康的疫苗id
	VaccineId string `json:"vaccine_id,omitempty" xml:"vaccine_id,omitempty"`
	// 阿里健康的pov id
	PovId string `json:"pov_id,omitempty" xml:"pov_id,omitempty"`
}
