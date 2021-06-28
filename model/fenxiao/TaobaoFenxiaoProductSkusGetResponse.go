package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
SKU查询接口 APIResponse
taobao.fenxiao.product.skus.get

产品sku查询
*/
type TaobaoFenxiaoProductSkusGetAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductSkusGetResponse
}

type TaobaoFenxiaoProductSkusGetResponse struct {
    XMLName xml.Name `xml:"fenxiao_product_skus_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // sku信息
    
    Skus   []FenxiaoSku `json:"skus,omitempty" xml:"skus>fenxiao_sku,omitempty"`
    
    
    // 记录数量
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
