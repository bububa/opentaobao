package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
零售云小程序获取登录用户信息 APIResponse
alibaba.lsy.miniapp.user.get

零售云小程序，通过授权码获取登录的卖家账号信息
*/
type AlibabaLsyMiniappUserGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaLsyMiniappUserGetResponse `json:"alibaba_lsy_miniapp_user_get_response,omitempty"`
}

type AlibabaLsyMiniappUserGetResponse struct {

    // 响应内容
    Result   *MiniAppResult `json:"result,omitempty"`

}
