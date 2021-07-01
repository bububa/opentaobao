package normalvisa

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
代填办理人信息 API返回值 
taobao.alitrip.travel.normalvisa.storeuser

卖家代填买家填写办理人信息
*/
type TaobaoAlitripTravelNormalvisaStoreuserAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelNormalvisaStoreuserAPIResponseModel
}

// 代填办理人信息 成功返回结果
type TaobaoAlitripTravelNormalvisaStoreuserAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_travel_normalvisa_storeuser_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果：包含results数组代表成功
    Result   *TaobaoAlitripTravelNormalvisaStoreuserResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
