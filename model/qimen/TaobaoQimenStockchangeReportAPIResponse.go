package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstockchangereportAPIResponse 库存异动通知接口 API返回值
// taobao.qimen.stockchange.report
//
// taobao.qimen.stockchange.report
type TaobaoqimenstockchangereportAPIResponse struct {
	model.CommonResponse
	TaobaoqimenstockchangereportAPIResponseModel
}

// TaobaoqimenstockchangereportAPIResponseModel is 库存异动通知接口 成功返回结果
type TaobaoqimenstockchangereportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_stockchange_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimenstockchangereportResponse `json:"response,omitempty" xml:"response,omitempty"`
}
