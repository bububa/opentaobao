package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取用户开通的topic列表 APIResponse
taobao.tmc.user.topics.get

获取用户开通的topic列表，授权消息使用
*/
type TaobaoTmcUserTopicsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTmcUserTopicsGetResponse `json:"taobao_tmc_user_topics_get_response,omitempty"`
}

type TaobaoTmcUserTopicsGetResponse struct {

    // 错误信息
    ResultMessage   string `json:"result_message,omitempty"`

    // topic列表
    Topics   []String `json:"topics,omitempty"`

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

}
