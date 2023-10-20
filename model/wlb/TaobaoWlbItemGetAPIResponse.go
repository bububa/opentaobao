package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemGetAPIResponse 根据商品ID获取商品信息 API返回值
// taobao.wlb.item.get
//
// 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
type TaobaoWlbItemGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbItemGetAPIResponseModel
}

// TaobaoWlbItemGetAPIResponseModel is 根据商品ID获取商品信息 成功返回结果
type TaobaoWlbItemGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品信息
	Item *WlbItem `json:"item,omitempty" xml:"item,omitempty"`
}
