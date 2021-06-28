package jae

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
活动平台弱登录接口 APIResponse
taobao.aplatform.weakget

无线活动平台的开放接口，提供商品信息等的读操作
*/
type TaobaoAplatformWeakgetAPIResponse struct {
    model.CommonResponse
    TaobaoAplatformWeakgetResponse
}

type TaobaoAplatformWeakgetResponse struct {
    XMLName xml.Name `xml:"aplatform_weakget_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoAplatformWeakgetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
