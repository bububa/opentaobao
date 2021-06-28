package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户开通的topic列表 APIResponse
taobao.tmc.user.topics.get

获取用户开通的topic列表，授权消息使用
*/
type TaobaoTmcUserTopicsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmc_user_topics_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误信息
    
    ResultMessage   string `json:"result_message,omitempty" xml:"