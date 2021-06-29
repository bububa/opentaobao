package normalvisa

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代填办理人信息 APIResponse
taobao.alitrip.travel.normalvisa.storeuser

卖家代填买家填写办理人信息
*/
type TaobaoAlitripTravelNormalvisaStoreuserAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelNormalvisaStoreuserResponse
}

type TaobaoAlitripTravelNormalvisaStoreuserResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_normalvisa_storeuser_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果：包含results数组代表成功
    
    Result   *TaobaoAlitripTravelNormalvisaStoreuserResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
