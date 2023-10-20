package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorySynchronizeReportAPIResponse 库存状态同步确认接口 API返回值
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
type TaobaoQimenInventorySynchronizeReportAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventorySynchronizeReportAPIResponseModel
}

// TaobaoQimenInventorySynchronizeReportAPIResponseModel is 库存状态同步确认接口 成功返回结果
type TaobaoQimenInventorySynchronizeReportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventory_synchronize_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenInventorySynchronizeReportResponse `json:"response,omitempty" xml:"response,omitempty"`
}
