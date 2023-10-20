package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventorysynchronizereportAPIResponse 库存状态同步确认接口 API返回值
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
type TaobaoqimeninventorysynchronizereportAPIResponse struct {
	model.CommonResponse
	TaobaoqimeninventorysynchronizereportAPIResponseModel
}

// TaobaoqimeninventorysynchronizereportAPIResponseModel is 库存状态同步确认接口 成功返回结果
type TaobaoqimeninventorysynchronizereportAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventory_synchronize_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimeninventorysynchronizereportResponse `json:"response,omitempty" xml:"response,omitempty"`
}
