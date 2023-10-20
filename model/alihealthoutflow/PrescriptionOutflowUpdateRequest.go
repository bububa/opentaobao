package alihealthoutflow

import (
	"sync"
)

// PrescriptionOutflowUpdateRequest 结构体
type PrescriptionOutflowUpdateRequest struct {
	// 诊断(非空)
	Diagnoses []Diagnose `json:"diagnoses,omitempty" xml:"diagnoses>diagnose,omitempty"`
	// 药品
	Drugs []Drugs `json:"drugs,omitempty" xml:"drugs>drugs,omitempty"`
	// 患者姓名(非空)
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
	// 操作人或患者手机号-用于接收核销码短信(非空)
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 操作人支付宝userId(可空)
	AlipayUserId string `json:"alipay_user_id,omitempty" xml:"alipay_user_id,omitempty"`
	// 患者年龄(非空)
	Age string `json:"age,omitempty" xml:"age,omitempty"`
	// 患者性别(非空)
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
	// 患者现住址(可空)
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 就诊号(可空)
	VisitId string `json:"visit_id,omitempty" xml:"visit_id,omitempty"`
	// 就诊类型（首诊、复诊）(可空)
	VisitType string `json:"visit_type,omitempty" xml:"visit_type,omitempty"`
	// 渠道、医院(非空)
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 就诊科室id(可空)
	DetpId string `json:"detp_id,omitempty" xml:"detp_id,omitempty"`
	// 就诊科室名称(非空)
	DeptName string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	// 就诊医生id(可空)
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
	// 就诊医生姓名(非空)
	DoctorName string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
	// 患者id(可空)
	PatientId string `json:"patient_id,omitempty" xml:"patient_id,omitempty"`
	// 就诊时间(可空)
	VisitTime string `json:"visit_time,omitempty" xml:"visit_time,omitempty"`
	// 主诉(可空)
	MainTell string `json:"main_tell,omitempty" xml:"main_tell,omitempty"`
	// 现病史(可空)
	ProblemNow string `json:"problem_now,omitempty" xml:"problem_now,omitempty"`
	// 一般检查(可空)
	BodyCheck string `json:"body_check,omitempty" xml:"body_check,omitempty"`
	// 医生嘱言(可空)
	DoctorAdvice string `json:"doctor_advice,omitempty" xml:"doctor_advice,omitempty"`
	// 处方编号(非空)
	RxNo string `json:"rx_no,omitempty" xml:"rx_no,omitempty"`
	// 处方类型(可空) COMMON-普通处方（默认） CHILDREN-儿童处方
	RxType string `json:"rx_type,omitempty" xml:"rx_type,omitempty"`
	// 药品类型(可空)WEST-西药（默认） TCM-草药
	DrugType string `json:"drug_type,omitempty" xml:"drug_type,omitempty"`
	// 费用类型(可空)OWN_EXPENSE-自费（默认）MEDICAL_INSURANCE-医保
	FeeType string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// 处方流转医保备案编号(医保支付时必填)
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 浙江省平台(可空)
	PlatformCode string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
	// 扩展属性JSON
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 身份证号(可空)
	IdCard string `json:"id_card,omitempty" xml:"id_card,omitempty"`
	// 医院档案类型(可空)VISIT_CARD就诊卡、ID_CARD身份证、MEDICAL_INSURANCE医保卡、MEDICAL_INSURANCE_ELECTRONIC_VOUCHER医保电子凭证、HOSPITAL_MEDICAL_ID医院病历号
	ArchivesType string `json:"archives_type,omitempty" xml:"archives_type,omitempty"`
	// 处方来源(可空)INTERNET_HOSPITAL_PRESCRIPTION互联网医院处方外配、DEPART_PRESCRIPTION门诊处方外配、CONSUMABLES_ADVICE耗材医嘱流转
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 模板(可空)INTERNET_HOSPITAL_PRESCRIPTION互联网医院处方笺、ELECTRONIC_PRESCRIPTION电子处方笺、DOCTOR_ADVICE医嘱单
	Template string `json:"template,omitempty" xml:"template,omitempty"`
	// 患者就医医院的档案/卡号ID
	ArchivesId string `json:"archives_id,omitempty" xml:"archives_id,omitempty"`
	// 患者类型(可空)：OWN_EXPENSE-自费、PROVINCE_MEDICAL_INSURANCE-省医保、CITY_MEDICAL_INSURANCE-市医保、BUSINESS_INSURANCE-商保
	PatientType string `json:"patient_type,omitempty" xml:"patient_type,omitempty"`
	// 患者参保地(可空)：杭州市本级、衢州市本级、浙江省本级、。。。。
	PatientInsuredRegion string `json:"patient_insured_region,omitempty" xml:"patient_insured_region,omitempty"`
	// 医保结算发生地(可空)：医保中心代码（行政区域代码）
	InsuranceSettlementRegion string `json:"insurance_settlement_region,omitempty" xml:"insurance_settlement_region,omitempty"`
	// 同步his错误信息(可空)
	SyncHisErrMsg string `json:"sync_his_err_msg,omitempty" xml:"sync_his_err_msg,omitempty"`
	// 同步his结果(非空)
	SyncHisResult bool `json:"sync_his_result,omitempty" xml:"sync_his_result,omitempty"`
}

var poolPrescriptionOutflowUpdateRequest = sync.Pool{
	New: func() any {
		return new(PrescriptionOutflowUpdateRequest)
	},
}

// GetPrescriptionOutflowUpdateRequest() 从对象池中获取PrescriptionOutflowUpdateRequest
func GetPrescriptionOutflowUpdateRequest() *PrescriptionOutflowUpdateRequest {
	return poolPrescriptionOutflowUpdateRequest.Get().(*PrescriptionOutflowUpdateRequest)
}

// ReleasePrescriptionOutflowUpdateRequest 释放PrescriptionOutflowUpdateRequest
func ReleasePrescriptionOutflowUpdateRequest(v *PrescriptionOutflowUpdateRequest) {
	v.Diagnoses = v.Diagnoses[:0]
	v.Drugs = v.Drugs[:0]
	v.PatientName = ""
	v.MobilePhone = ""
	v.AlipayUserId = ""
	v.Age = ""
	v.Sex = ""
	v.Address = ""
	v.VisitId = ""
	v.VisitType = ""
	v.ChannelCode = ""
	v.DetpId = ""
	v.DeptName = ""
	v.DoctorId = ""
	v.DoctorName = ""
	v.PatientId = ""
	v.VisitTime = ""
	v.MainTell = ""
	v.ProblemNow = ""
	v.BodyCheck = ""
	v.DoctorAdvice = ""
	v.RxNo = ""
	v.RxType = ""
	v.DrugType = ""
	v.FeeType = ""
	v.CardNumber = ""
	v.PlatformCode = ""
	v.Attributes = ""
	v.IdCard = ""
	v.ArchivesType = ""
	v.Source = ""
	v.Template = ""
	v.ArchivesId = ""
	v.PatientType = ""
	v.PatientInsuredRegion = ""
	v.InsuranceSettlementRegion = ""
	v.SyncHisErrMsg = ""
	v.SyncHisResult = false
	poolPrescriptionOutflowUpdateRequest.Put(v)
}
