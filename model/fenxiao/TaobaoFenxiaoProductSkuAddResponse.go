package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品sku添加接口 APIResponse
taobao.fenxiao.product.sku.add

添加产品SKU信息
*/
type TaobaoFenxiaoProductSkuAddAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductSkuAddResponse
}

type TaobaoFenxiaoProductSkuAddResponse struct {
    XMLName xml.Name `xml:"fenxiao_product_sku_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 操作时间
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`

    
}
