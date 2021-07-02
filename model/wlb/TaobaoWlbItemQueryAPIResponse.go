package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbItemQueryAPIResponse 分页查询商品 API返回值
// taobao.wlb.item.query
//
// 根据状态、卖家、SKU等信息查询商品列表
type TaobaoWlbItemQueryAPIResponse struct {
	model.CommonResponse
	TaobaoWlbItemQueryAPIResponseModel
}

// TaobaoWlbItemQueryAPIResponseModel is 分页查询商品 成功返回结果
type TaobaoWlbItemQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 商品信息列表
	ItemList []WlbItem `json:"item_list,omitempty" xml:"item_list>wlb_item,omitempty"`
}
