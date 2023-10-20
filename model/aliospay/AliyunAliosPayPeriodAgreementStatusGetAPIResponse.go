package aliospay

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunaliospayperiodagreementstatusgetAPIResponse 查询周期扣款签约状态 API返回值
// aliyun.alios.pay.period.agreement.status.get
//
// 查询周期扣款签约状态
type AliyunaliospayperiodagreementstatusgetAPIResponse struct {
	model.CommonResponse
	AliyunaliospayperiodagreementstatusgetAPIResponseModel
}

// AliyunaliospayperiodagreementstatusgetAPIResponseModel is 查询周期扣款签约状态 成功返回结果
type AliyunaliospayperiodagreementstatusgetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alios_pay_period_agreement_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应参数
	AliospayResponse *AliOspayResponse `json:"aliospay_response,omitempty" xml:"aliospay_response,omitempty"`
}
