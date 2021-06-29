package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商查询产品信息 APIResponse
taobao.fenxiao.distributor.products.get

分销商查询供应商产品信息
*/
type TaobaoFenxiaoDistributorProductsGetAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoDistributorProductsGetResponse
}

type TaobaoFenxiaoDistributorProductsGetResponse struct {
    XMLName xml.Name `xml:"fenxiao_distributor_products_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否存在下一页
    
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`

    
    // 产品对象记录集。返回 FenxiaoProduct 包含的字段信息。
    
    Products   []FenxiaoProduct `json:"products,omitempty" xml:"products>fenxiao_product,omitempty"`
    
    
}
