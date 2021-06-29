package normalvisa

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取物流公司信息 APIResponse
taobao.alitrip.travel.normalvisa.getcompany

获取物流公司信息
*/
type TaobaoAlitripTravelNormalvisaGetcompanyAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelNormalvisaGetcompanyResponse
}

type TaobaoAlitripTravelNormalvisaGetcompanyResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_normalvisa_getcompany_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果，有返回代表成功
    
    Result   *TaobaoAlitripTravelNormalvisaGetcompanyResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
