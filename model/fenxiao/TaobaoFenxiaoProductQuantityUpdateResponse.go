package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
产品库存修改 API返回值 
taobao.fenxiao.product.quantity.update

修改产品库存信息，支持全量修改以及增量修改两种方式
*/
type TaobaoFenxiaoProductQuantityUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductQuantityUpdateResponse
}

// 产品库存修改 成功返回结果
type TaobaoFenxiaoProductQuantityUpdateResponse struct {
    XMLName xml.Name `xml:"fenxiao_product_quantity_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 操作结果
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    // 操作时间
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
}
