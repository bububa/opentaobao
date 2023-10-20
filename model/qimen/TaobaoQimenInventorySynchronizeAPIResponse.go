package qimen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorySynchronizeAPIResponse 库存状态同步接口 API返回值
// taobao.qimen.inventory.synchronize
//
// ERP通过该接口同步指定商品的库存信息
type TaobaoQimenInventorySynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenInventorySynchronizeAPIResponseModel
}

// TaobaoQimenInventorySynchronizeAPIResponseModel is 库存状态同步接口 成功返回结果
type TaobaoQimenInventorySynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_inventory_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenInventorySynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}
