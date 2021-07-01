package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenOrderSnReportAPIResponse
订单SN通知接口 API返回值
taobao.qimen.order.sn.report

WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表 */
type TaobaoQimenOrderSnReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenOrderSnReportAPIResponseModel
}

// TaobaoQimenOrderSnReportAPIResponseModel is 订单SN通知接口 成功返回结果
type TaobaoQimenOrderSnReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_order_sn_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenOrderSnReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}
