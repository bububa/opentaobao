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
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_product_sku_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 操作结果
    
    Result   bool `json:"result,omitempty" xml:"