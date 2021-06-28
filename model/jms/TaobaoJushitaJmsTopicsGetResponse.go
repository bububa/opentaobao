package jms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据用户nick获取开通的消息列表 APIResponse
taobao.jushita.jms.topics.get

根据用户nick获取开通的消息列表
*/
type TaobaoJushitaJmsTopicsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoJushitaJmsTopicsGetResponse `json:"jushita_jms_topics_get_response,omitempty"` 
    TaobaoJushitaJmsTopicsGetResponse
}

/* model for simplify = false
type TaobaoJushitaJmsTopicsGetResponse struct {

    // 错误信息
    
    ResultMessage   string `json:"result_message,omitempty"`
    

    // topic列表
    
    Results  struct {
        String  []string `json:"string,omitempty"`
    } `json:"results,omitempty"`
    

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

}
*/

type TaobaoJushitaJmsTopicsGetResponse struct {

    // 错误信息
    ResultMessage   string `json:"result_message,omitempty"`

    // topic列表
    Results   []string `json:"results,omitempty"`

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

}
