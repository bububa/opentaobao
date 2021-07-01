package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbItemMapGetAPIResponse
根据物流宝商品ID查询商品映射关系 API返回值
taobao.wlb.item.map.get

根据物流宝商品ID查询商品映射关系 */
type TaobaoWlbItemMapGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbItemMapGetAPIResponseModel
}

// TaobaoWlbItemMapGetAPIResponseModel is 根据物流宝商品ID查询商品映射关系 成功返回结果
type TaobaoWlbItemMapGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_map_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 外部商品实体
	OutEntityItemList []OutEntityItem `json:"out_entity_item_list,omitempty" xml:"out_entity_item_list>out_entity_item,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
