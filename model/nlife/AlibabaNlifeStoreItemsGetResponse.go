package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取门店的商品列表(在售|已下架|全部) API返回值 
alibaba.nlife.store.items.get

利用该接口可以获取到零售+商品服务中符合条件的商品列表，包括在售的、已下架的或者是所有状态的商品。
*/
type AlibabaNlifeStoreItemsGetAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeStoreItemsGetResponse
}

// 获取门店的商品列表(在售|已下架|全部) 成功返回结果
type AlibabaNlifeStoreItemsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_store_items_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 搜索到符合条件的结果总数
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
    // 具体的商品信息列表
    ItemList   []RetailItemTopDO `json:"item_list,omitempty" xml:"item_list>retail_item_top_do,omitempty"`
    // true-查询成功;false-查询失败
    Succes   bool `json:"succes,omitempty" xml:"succes,omitempty"`
}
