package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemAddAPIResponse 添加单个物流宝商品 API返回值
// taobao.wlb.item.add
//
// 添加物流宝商品，支持物流宝子商品和属性添加
type TaobaoWlbItemAddAPIResponse struct {
	model.CommonResponse
	TaobaoWlbItemAddAPIResponseModel
}

// TaobaoWlbItemAddAPIResponseModel is 添加单个物流宝商品 成功返回结果
type TaobaoWlbItemAddAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 新增的商品
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
