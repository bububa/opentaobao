package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品sku编辑接口 APIResponse
taobao.fenxiao.product.sku.update

产品SKU信息更新
*/
type TaobaoFenxiaoProductSkuUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductSkuUpdateResponse
}

type TaobaoFenxiaoProductSkuUpdateResponse struct {
    XMLName xml.Name `xml:"fenxiao_product_sku_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 操作结果
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 操作时间
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`

    
}
