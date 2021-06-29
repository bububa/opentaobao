package fivee

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
国产商品资质查询 APIResponse
taobao.fivee.innerproduct.get

资质共享平台，国产商品查询
*/
type TaobaoFiveeInnerproductGetAPIResponse struct {
    model.CommonResponse
    TaobaoFiveeInnerproductGetResponse
}

type TaobaoFiveeInnerproductGetResponse struct {
    XMLName xml.Name `xml:"fivee_innerproduct_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoFiveeInnerproductGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
