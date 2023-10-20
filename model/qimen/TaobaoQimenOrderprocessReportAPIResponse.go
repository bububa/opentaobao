package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenorderprocessreportAPIResponse 订单流水通知接口 API返回值
// taobao.qimen.orderprocess.report
//
// taobao.qimen.orderprocess.report
type TaobaoqimenorderprocessreportAPIResponse struct {
	model.CommonResponse
	TaobaoqimenorderprocessreportAPIResponseModel
}

// TaobaoqimenorderprocessreportAPIResponseModel is 订单流水通知接口 成功返回结果
type TaobaoqimenorderprocessreportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_orderprocess_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenorderprocessreportResponse `json:"response,omitempty" xml:"response,omitempty"`
}
