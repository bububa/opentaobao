package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查找IC商品或分销商品与后端商品的关联信息 APIResponse
taobao.scitem.map.query

查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku
*/
type TaobaoScitemMapQueryAPIResponse struct {
    model.CommonResponse
    TaobaoScitemMapQueryResponse
}

type TaobaoScitemMapQueryResponse struct {
    XMLName xml.Name `xml:"scitem_map_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 后端商品映射列表
    
    ScItemMaps   []ScItemMap `json:"sc_item_maps,omitempty" xml:"sc_item_maps>sc_item_map,omitempty"`
    
    
}
