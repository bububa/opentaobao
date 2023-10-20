package logistic

import (
	"sync"
)

// TmsCollectRequest 结构体
type TmsCollectRequest struct {
	// 电联信息
	PhoneCallInfos []TmsPhoneCallInfoDto `json:"phone_call_infos,omitempty" xml:"phone_call_infos>tms_phone_call_info_dto,omitempty"`
	// 其他作业信息明细
	ExtendOperateInfos []TmsExtendOperateInfosDto `json:"extend_operate_infos,omitempty" xml:"extend_operate_infos>tms_extend_operate_infos_dto,omitempty"`
	// 业务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 运单号（运单号及包裹重量至少有一个）
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 验视照片
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 单据标识
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 包裹长度 (单位：cm)，小数点后2位
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 包裹宽度（单位：cm)，小数点后2位
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 业务服务
	ServiceFlag string `json:"service_flag,omitempty" xml:"service_flag,omitempty"`
	// 实际重量（取件称重重量），单位g ，小数点后2位
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 资源编码
	TmsCpCode string `json:"tms_cp_code,omitempty" xml:"tms_cp_code,omitempty"`
	// 预估运费（单位：分）
	EstimateFee string `json:"estimate_fee,omitempty" xml:"estimate_fee,omitempty"`
	// 计费重量，单位g，小数点后2位
	ChargedWeight string `json:"charged_weight,omitempty" xml:"charged_weight,omitempty"`
	// 包裹高度 (单位：cm)，小数点后2位
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 服务商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 配资源编码
	DeliveryCode string `json:"delivery_code,omitempty" xml:"delivery_code,omitempty"`
	// 计抛重量，单位g，小数点后2位。条件必填，计抛重量与包裹长宽高信息需要同时传入
	ThrowingWeight string `json:"throwing_weight,omitempty" xml:"throwing_weight,omitempty"`
	// 0-未核实 1-已核实
	WeightIsVerified string `json:"weight_is_verified,omitempty" xml:"weight_is_verified,omitempty"`
	// 扩展字段
	Feature *TmsCollectFeatureDto `json:"feature,omitempty" xml:"feature,omitempty"`
}

var poolTmsCollectRequest = sync.Pool{
	New: func() any {
		return new(TmsCollectRequest)
	},
}

// GetTmsCollectRequest() 从对象池中获取TmsCollectRequest
func GetTmsCollectRequest() *TmsCollectRequest {
	return poolTmsCollectRequest.Get().(*TmsCollectRequest)
}

// ReleaseTmsCollectRequest 释放TmsCollectRequest
func ReleaseTmsCollectRequest(v *TmsCollectRequest) {
	v.PhoneCallInfos = v.PhoneCallInfos[:0]
	v.ExtendOperateInfos = v.ExtendOperateInfos[:0]
	v.ServiceType = ""
	v.MailNo = ""
	v.PicUrl = ""
	v.BizCode = ""
	v.Length = ""
	v.Width = ""
	v.ServiceFlag = ""
	v.Weight = ""
	v.TmsCpCode = ""
	v.EstimateFee = ""
	v.ChargedWeight = ""
	v.Height = ""
	v.SupplierId = ""
	v.DeliveryCode = ""
	v.ThrowingWeight = ""
	v.WeightIsVerified = ""
	v.Feature = nil
	poolTmsCollectRequest.Put(v)
}
