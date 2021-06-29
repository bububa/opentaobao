package normalvisa

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新签证办理进度 APIResponse
taobao.alitrip.travel.normalvisa.updatepersonstauts

更新签证办理进度
*/
type TaobaoAlitripTravelNormalvisaUpdatepersonstautsAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelNormalvisaUpdatepersonstautsResponse
}

type TaobaoAlitripTravelNormalvisaUpdatepersonstautsResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_normalvisa_updatepersonstauts_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果：包含result更新成功
    
    Result   *TaobaoAlitripTravelNormalvisaUpdatepersonstautsResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
