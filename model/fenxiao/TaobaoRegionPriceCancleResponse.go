package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消区域价格 APIResponse
taobao.region.price.cancle

取消区域价格
*/
type TaobaoRegionPriceCancleAPIResponse struct {
    model.CommonResponse
    TaobaoRegionPriceCancleResponse
}

type TaobaoRegionPriceCancleResponse struct {
    XMLName xml.Name `xml:"region_price_cancle_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
