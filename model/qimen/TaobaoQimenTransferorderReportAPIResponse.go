package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTransferorderReportAPIResponse
调拨单通知 API返回值
taobao.qimen.transferorder.report

调拨单通知 */
type TaobaoQimenTransferorderReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenTransferorderReportAPIResponseModel
}

// TaobaoQimenTransferorderReportAPIResponseModel is 调拨单通知 成功返回结果
type TaobaoQimenTransferorderReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_transferorder_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenTransferorderReportStruct `json:"response,omitempty" xml:"response,omitempty"`
}
