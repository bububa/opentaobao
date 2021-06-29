package normalvisa

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取签证记录 APIResponse
taobao.alitrip.travel.normalvisa.get

用于获取普通签证的记录信息
*/
type TaobaoAlitripTravelNormalvisaGetAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelNormalvisaGetResponse
}

type TaobaoAlitripTravelNormalvisaGetResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_normalvisa_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TaobaoAlitripTravelNormalvisaGetResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
