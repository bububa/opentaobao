package aliospay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayRecordListAPIResponse 支付记录批量查询接口 API返回值
// aliyun.alios.pay.record.list
//
// 商户用来对账的接口
type AliyunAliosPayRecordListAPIResponse struct {
	model.CommonResponse
	AliyunAliosPayRecordListAPIResponseModel
}

// AliyunAliosPayRecordListAPIResponseModel is 支付记录批量查询接口 成功返回结果
type AliyunAliosPayRecordListAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_record_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOSPayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}
