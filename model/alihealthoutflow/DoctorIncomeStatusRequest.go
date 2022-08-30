package alihealthoutflow

// DoctorIncomeStatusRequest 结构体
type DoctorIncomeStatusRequest struct {
	// 说明信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 医生id
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 是否打款成功
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 医生收入流水号
	IncomeId int64 `json:"income_id,omitempty" xml:"income_id,omitempty"`
}
