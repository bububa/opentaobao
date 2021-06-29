package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AG物流签收状态写接口 API返回值 
taobao.nextone.logistics.sign.update

商家上传退货的签收状态给AG
*/
type TaobaoNextoneLogisticsSignUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoNextoneLogisticsSignUpdateResponse
}

// AG物流签收状态写接口 成功返回结果
type TaobaoNextoneLogisticsSignUpdateResponse struct {
    XMLName xml.Name `xml:"nextone_logistics_sign_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *TaobaoNextoneLogisticsSignUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
