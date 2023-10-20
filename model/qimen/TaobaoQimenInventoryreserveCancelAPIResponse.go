package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryreserveCancelAPIResponse 库存预占取消接口 API返回值
// taobao.qimen.inventoryreserve.cancel
//
// 库存预占取消
type TaobaoQimenInventoryreserveCancelAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventoryreserveCancelAPIResponseModel
}

// TaobaoQimenInventoryreserveCancelAPIResponseModel is 库存预占取消接口 成功返回结果
type TaobaoQimenInventoryreserveCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventoryreserve_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenInventoryreserveCancelResponse `json:"response,omitempty" xml:"response,omitempty"`
}
