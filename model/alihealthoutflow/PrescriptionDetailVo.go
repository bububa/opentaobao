package alihealthoutflow

// PrescriptionDetailVo 结构体
type PrescriptionDetailVo struct {
	// 病历报告单
	RevisitPicUrlList []string `json:"revisit_pic_url_list,omitempty" xml:"revisit_pic_url_list>string,omitempty"`
	// 药品
	DrugList []DrugDto `json:"drug_list,omitempty" xml:"drug_list>drug_dto,omitempty"`
	// 诊断
	DiagnoseList []DiagnoseVo `json:"diagnose_list,omitempty" xml:"diagnose_list>diagnose_vo,omitempty"`
	// 患者年龄
	PatientAge string `json:"patient_age,omitempty" xml:"patient_age,omitempty"`
	// 患者电话
	PatientTel string `json:"patient_tel,omitempty" xml:"patient_tel,omitempty"`
	// 审核药师姓名
	AuditPharmacistName string `json:"audit_pharmacist_name,omitempty" xml:"audit_pharmacist_name,omitempty"`
	// 科室名
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 省份证
	IdCard string `json:"id_card,omitempty" xml:"id_card,omitempty"`
	// 处方图地址
	PrescriptionPicUrl string `json:"prescription_pic_url,omitempty" xml:"prescription_pic_url,omitempty"`
	// 肝功能
	GanGongDetail string `json:"gan_gong_detail,omitempty" xml:"gan_gong_detail,omitempty"`
	// 既往史
	ProblemHistory string `json:"problem_history,omitempty" xml:"problem_history,omitempty"`
	// 淘宝订单号
	TaobaoOrderNo string `json:"taobao_order_no,omitempty" xml:"taobao_order_no,omitempty"`
	// 患者姓名
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
	// 配药药师姓名
	DispensingPharmacistName string `json:"dispensing_pharmacist_name,omitempty" xml:"dispensing_pharmacist_name,omitempty"`
	// 患者性别
	PatientSex string `json:"patient_sex,omitempty" xml:"patient_sex,omitempty"`
	// 医院名
	HospitalName string `json:"hospital_name,omitempty" xml:"hospital_name,omitempty"`
	// 肾功能
	ShenGongDetail string `json:"shen_gong_detail,omitempty" xml:"shen_gong_detail,omitempty"`
	// 处方号
	RxNo string `json:"rx_no,omitempty" xml:"rx_no,omitempty"`
	// 病历号
	DiseaseRecordId string `json:"disease_record_id,omitempty" xml:"disease_record_id,omitempty"`
	// 过敏
	GuoMinDetail string `json:"guo_min_detail,omitempty" xml:"guo_min_detail,omitempty"`
	// 处方药有效期（小时）
	EffectiveTime string `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
	// 医生姓名
	DoctorName string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
	// 医院id
	HospitalId string `json:"hospital_id,omitempty" xml:"hospital_id,omitempty"`
	// 出生年月日
	PatientBirthday string `json:"patient_birthday,omitempty" xml:"patient_birthday,omitempty"`
	// 处方审核时间
	PrescriptionAuditTime int64 `json:"prescription_audit_time,omitempty" xml:"prescription_audit_time,omitempty"`
	// 处开具时间
	PrescriptionCreateTime int64 `json:"prescription_create_time,omitempty" xml:"prescription_create_time,omitempty"`
	// 接诊结束时间
	EndReceiveTime int64 `json:"end_receive_time,omitempty" xml:"end_receive_time,omitempty"`
	// 开始接诊时间
	StartReceiveTime int64 `json:"start_receive_time,omitempty" xml:"start_receive_time,omitempty"`
}
