package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消区域价格 API返回值 
taobao.region.price.cancle

取消区域价格
*/
type TaobaoRegionPriceCancleAPIResponse struct {
    model.CommonResponse
    TaobaoRegionPriceCancleResponse
}

// 取消区域价格 成功返回结果
type TaobaoRegionPriceCancleResponse struct {
    XMLName xml.Name `xml:"region_price_cancle_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
