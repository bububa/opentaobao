package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新产品 APIResponse
taobao.fenxiao.product.update

更新分销平台产品数据，不传更新数据返回失败<br><br/>1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br>
*/
type TaobaoFenxiaoProductUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoFenxiaoProductUpdateResponse
}

type TaobaoFenxiaoProductUpdateResponse struct {
    XMLName xml.Name `xml:"fenxiao_product_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 产品ID
    
    Pid   int64 `json:"pid,omitempty" xml:"pid,omitempty"`

    
    // 更新时间，时间格式：yyyy-MM-dd HH:mm:ss
    
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`

    
}
