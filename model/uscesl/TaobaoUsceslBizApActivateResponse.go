package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
激活AP价签通讯模块 APIResponse
taobao.uscesl.biz.ap.activate

激活AP价签通讯模块
*/
type TaobaoUsceslBizApActivateAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizApActivateResponse
}

type TaobaoUsceslBizApActivateResponse struct {
    XMLName xml.Name `xml:"uscesl_biz_ap_activate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功与否看result.success，返回true或者false
    
    Result   *TaobaoUsceslBizApActivateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
