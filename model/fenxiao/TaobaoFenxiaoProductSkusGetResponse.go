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
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_product_skus_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // sku信息
    
    Skus   []FenxiaoSku `json:"skus,omitempty" xml:"