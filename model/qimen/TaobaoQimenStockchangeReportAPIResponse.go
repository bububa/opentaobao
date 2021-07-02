package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStockchangeReportAPIResponse 库存异动通知接口 API返回值
// taobao.qimen.stockchange.report
//
// WMS调用奇门的接口,将库存异动信息信息回传给ERP
type TaobaoQimenStockchangeReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenStockchangeReportAPIResponseModel
}

// TaobaoQimenStockchangeReportAPIResponseModel is 库存异动通知接口 成功返回结果
type TaobaoQimenStockchangeReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_stockchange_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenStockchangeReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}
