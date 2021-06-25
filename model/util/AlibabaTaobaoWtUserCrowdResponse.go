package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
校验用户是否为新人 APIResponse
alibaba.taobao.wt.user.crowd

增加isv接口
根据入参手机号判断是否为新人类型
*/
type AlibabaTaobaoWtUserCrowdAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTaobaoWtUserCrowdResponse `json:"alibaba_taobao_wt_user_crowd_response,omitempty"`
}

type AlibabaTaobaoWtUserCrowdResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
