package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventoryreservecancelAPIResponse 库存预占取消接口 API返回值
// taobao.qimen.inventoryreserve.cancel
//
// 库存预占取消
type TaobaoqimeninventoryreservecancelAPIResponse struct {
	model.CommonResponse
	TaobaoqimeninventoryreservecancelAPIResponseModel
}

// TaobaoqimeninventoryreservecancelAPIResponseModel is 库存预占取消接口 成功返回结果
type TaobaoqimeninventoryreservecancelAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventoryreserve_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoqimeninventoryreservecancelResponse `json:"response,omitempty" xml:"response,omitempty"`
}
