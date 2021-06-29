package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品的详情信息 APIResponse
alibaba.nlife.store.itemdetail.get

查询零售加平台上单个商品的详情信息
*/
type AlibabaNlifeStoreItemdetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeStoreItemdetailGetResponse
}

type AlibabaNlifeStoreItemdetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_store_itemdetail_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // true-查询成功;false-查询失败
    
    Succes   bool `json:"succes,omitempty" xml:"succes,omitempty"`

    
    // 商品详情信息
    
    Item   *RetailItemTopDo `json:"item,omitempty" xml:"item,omitempty"`

    
}
