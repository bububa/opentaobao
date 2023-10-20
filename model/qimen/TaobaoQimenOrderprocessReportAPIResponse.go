package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderprocessReportAPIResponse 订单流水通知接口 API返回值
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
type TaobaoQimenOrderprocessReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenOrderprocessReportAPIResponseModel
}

// TaobaoQimenOrderprocessReportAPIResponseModel is 订单流水通知接口 成功返回结果
type TaobaoQimenOrderprocessReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_orderprocess_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenOrderprocessReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}
