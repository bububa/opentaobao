package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据outerCode查询商品 APIResponse
taobao.scitem.outercode.get

根据outerCode查询商品
*/
type TaobaoScitemOutercodeGetAPIResponse struct {
    model.CommonResponse
    TaobaoScitemOutercodeGetResponse
}

type TaobaoScitemOutercodeGetResponse struct {
    XMLName xml.Name `xml:"scitem_outercode_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 后台商品
    
    ScItem   *ScItem `json:"sc_item,omitempty" xml:"sc_item,omitempty"`

    
}
