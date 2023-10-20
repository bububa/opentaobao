package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemlackReportAPIResponse 发货单缺货通知接口 API返回值
// taobao.qimen.itemlack.report
//
// WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
type TaobaoQimenItemlackReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenItemlackReportAPIResponseModel
}

// TaobaoQimenItemlackReportAPIResponseModel is 发货单缺货通知接口 成功返回结果
type TaobaoQimenItemlackReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_itemlack_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenItemlackReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}
