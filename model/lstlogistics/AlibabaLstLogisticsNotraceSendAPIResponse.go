package lstlogistics

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstlogisticsnotracesendAPIResponse 供应商-异云-无需物流发货 API返回值
// alibaba.lst.logistics.notrace.send
//
// 异地云仓的订单，使用无需物流的发货方式来变更订单发货状态
type AlibabalstlogisticsnotracesendAPIResponse struct {
	model.CommonResponse
	AlibabalstlogisticsnotracesendAPIResponseModel
}

// AlibabalstlogisticsnotracesendAPIResponseModel is 供应商-异云-无需物流发货 成功返回结果
type AlibabalstlogisticsnotracesendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_logistics_notrace_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabalstlogisticsnotracesendResult `json:"result,omitempty" xml:"result,omitempty"`
}
