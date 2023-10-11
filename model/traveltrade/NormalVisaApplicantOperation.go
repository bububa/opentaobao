package traveltrade

// NormalVisaApplicantOperation 结构体
type NormalVisaApplicantOperation struct {
	// 签证需补充材料（即状态为1012）时必填，0:护照，1:证件照，2:申请表，3:身份证，4:户口本，5:暂住证，6:在职收入证明，7:营业执照，8:组织机构代码证，9:结婚证，10:个人信息处理同意书，11:退休证，12:保险订单，13:在读证明，14:机票预订证明，15:酒店预订证明，16:财力证明，17:房产证，18:汽车驾驶证，19:社保缴纳记录，20:学校准假证明，21:儿童出生医学证明，22:未成年人亲属关系证明，23:其他材料，24:银行存款证明，25:学生证，26:其他材料2，27:其他材料3，28:居住证，29:车辆登记证，34:保险声明，36:出行同意书，38:职业证明，39:以往申根签证页
	NeedModifyDocTypes []string `json:"need_modify_doc_types,omitempty" xml:"need_modify_doc_types>string,omitempty"`
	// 必填，申请人ID
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 可选，备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 代填申请人信息。字段注释：1.sex(性别),值:M/F;2.nationality(国籍),值:CHN(中国大陆),HKG(中国香港),MAC(中国澳门),USA(美国),CAN(加拿大)
	ApplicantFormDataJson string `json:"applicant_form_data_json,omitempty" xml:"applicant_form_data_json,omitempty"`
	// 特殊必填，上传该申请人 电子签结果。当该签证为电子签证且status值为1006（已收到签证结果）时 必填
	EtaInfo *NormalVisaETAInfo `json:"eta_info,omitempty" xml:"eta_info,omitempty"`
	// 特殊必填，上传该申请人 签证结果寄回物流信息。当status值为1013（已寄回结果）时，必填
	LogisticsInfo *NormalVisaLogisticsInfo `json:"logistics_info,omitempty" xml:"logistics_info,omitempty"`
	// 必填，申请人状态推进，本次操作需要推进到的目标状态。具体状态值枚举及推进流程详见：https://open.alitrip.com/docs/doc.htm?spm=a21tt.7629140.0.0.fYvMkZ&amp;docType=1&amp;articleId=108046&amp;previewCode=9D5F931C2254C7B3FE16B8DB7F9CECB4
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 特殊必填，上传该申请人 预约面试信息。当status值为1007（已预约面试）时 必填
	AppointmentInfo *NormalVisaAppointmentInfo `json:"appointment_info,omitempty" xml:"appointment_info,omitempty"`
}
