package alihealthoutflow

// AsyncPrescribeDetailRequest 结构体
type AsyncPrescribeDetailRequest struct {
	// 处方号
	RxNo string `json:"rx_no,omitempty" xml:"rx_no,omitempty"`
	// 医院id
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
}
