package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取零售加商品详情信息 API返回值 
alibaba.nlife.store.itemdetails.get

批量获取零售加平台上的商品详情信息
*/
type AlibabaNlifeStoreItemdetailsGetAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeStoreItemdetailsGetResponse
}

// 批量获取零售加商品详情信息 成功返回结果
type AlibabaNlifeStoreItemdetailsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_store_itemdetails_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商品详情列表
    ItemList   []RetailItemTopDO `json:"item_list,omitempty" xml:"item_list>retail_item_top_do,omitempty"`
    // true-查询成功;false-查询失败
    Succes   bool `json:"succes,omitempty" xml:"succes,omitempty"`
    // 总数
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
}
