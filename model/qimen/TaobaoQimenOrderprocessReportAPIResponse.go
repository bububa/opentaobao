package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenOrderprocessReportAPIResponse
订单流水通知接口 API返回值
taobao.qimen.orderprocess.report

WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。 */
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
