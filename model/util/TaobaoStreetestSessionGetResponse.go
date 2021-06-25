package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据获取压测用户的sessionKey APIResponse
taobao.streetest.session.get

根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
*/
type TaobaoStreetestSessionGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoStreetestSessionGetResponse `json:"taobao_streetest_session_get_response,omitempty"`
}

type TaobaoStreetestSessionGetResponse struct {

    // 压测账号对应的sessionKey
    StreeTestSessionKey   string `json:"stree_test_session_key,omitempty"`

}
