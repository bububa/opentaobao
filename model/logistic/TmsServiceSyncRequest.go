package logistic

import (
	"sync"
)

// TmsServiceSyncRequest 结构体
type TmsServiceSyncRequest struct {
	// 图片信息
	PictureInfoList []TmsPictureInfoRequest `json:"picture_info_list,omitempty" xml:"picture_info_list>tms_picture_info_request,omitempty"`
	// 短信信息
	SmsInfoList []TmsSmsInfoRequest `json:"sms_info_list,omitempty" xml:"sms_info_list>tms_sms_info_request,omitempty"`
	// 电联信息
	PhoneCallInfoList []TmsPhoneCallInfoRequest `json:"phone_call_info_list,omitempty" xml:"phone_call_info_list>tms_phone_call_info_request,omitempty"`
	// 服务类型（1-送货上门；2-放指定驿站）
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 配品牌
	TmsBrandCode string `json:"tms_brand_code,omitempty" xml:"tms_brand_code,omitempty"`
	// 物流轨迹描述信息
	LogisticsDetailDesc string `json:"logistics_detail_desc,omitempty" xml:"logistics_detail_desc,omitempty"`
	// 送货上门签收类型 1-本人签收 2-他人代签收 3-前台 4-家门口 5-水表箱等指定位置
	SendSignType string `json:"send_sign_type,omitempty" xml:"send_sign_type,omitempty"`
	// 服务单code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 对应状态的操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 履约结果 TO_HOME, 送货上门 TO_CABINET, 入库入柜； TO_CUSTOMER_LOCATION, 放置指定地点（未上门）
	DeliveryResult string `json:"delivery_result,omitempty" xml:"delivery_result,omitempty"`
	// 包裹运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 状态（1-已揽收；2-派送中；3-已电联；4-已签收；5-已取消；6-异常）
	ServiceState string `json:"service_state,omitempty" xml:"service_state,omitempty"`
	// 电联结果 NOT_CONNECT, 未接通;  TO_HOME, 送货上门; TO_CABINET, 需入库入柜
	PhoneCallResult string `json:"phone_call_result,omitempty" xml:"phone_call_result,omitempty"`
	// 快递公司资源编码
	TmsCpCode string `json:"tms_cp_code,omitempty" xml:"tms_cp_code,omitempty"`
	// 业务类型（1-送货上门 ）
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 上门签收描述
	DeliveryResultRemark string `json:"delivery_result_remark,omitempty" xml:"delivery_result_remark,omitempty"`
}

var poolTmsServiceSyncRequest = sync.Pool{
	New: func() any {
		return new(TmsServiceSyncRequest)
	},
}

// GetTmsServiceSyncRequest() 从对象池中获取TmsServiceSyncRequest
func GetTmsServiceSyncRequest() *TmsServiceSyncRequest {
	return poolTmsServiceSyncRequest.Get().(*TmsServiceSyncRequest)
}

// ReleaseTmsServiceSyncRequest 释放TmsServiceSyncRequest
func ReleaseTmsServiceSyncRequest(v *TmsServiceSyncRequest) {
	v.PictureInfoList = v.PictureInfoList[:0]
	v.SmsInfoList = v.SmsInfoList[:0]
	v.PhoneCallInfoList = v.PhoneCallInfoList[:0]
	v.ServiceType = ""
	v.TmsBrandCode = ""
	v.LogisticsDetailDesc = ""
	v.SendSignType = ""
	v.BizCode = ""
	v.OperateTime = ""
	v.DeliveryResult = ""
	v.MailNo = ""
	v.ServiceState = ""
	v.PhoneCallResult = ""
	v.TmsCpCode = ""
	v.BusinessType = ""
	v.DeliveryResultRemark = ""
	poolTmsServiceSyncRequest.Put(v)
}
