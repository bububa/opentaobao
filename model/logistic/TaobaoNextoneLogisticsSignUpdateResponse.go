package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AG物流签收状态写接口 APIResponse
taobao.nextone.logistics.sign.update

商家上传退货的签收状态给AG
*/
type TaobaoNextoneLogisticsSignUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"nextone_logistics_sign_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *TaobaoNextoneLogisticsSignUpdateResult `json:"result,omitempty" xml:"