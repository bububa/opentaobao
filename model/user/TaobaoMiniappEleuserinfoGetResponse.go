package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取饿了么用户信息 APIResponse
taobao.miniapp.eleuserinfo.get

获取饿了么用户信息
*/
type TaobaoMiniappEleuserinfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoMiniappEleuserinfoGetResponse `json:"taobao_miniapp_eleuserinfo_get_response,omitempty"`
}

type TaobaoMiniappEleuserinfoGetResponse struct {

    // traceId
    TraceId   string `json:"trace_id,omitempty"`

    // 1
    Result   *EleUicInfo `json:"result,omitempty"`

}
