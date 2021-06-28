package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
无线开放内容检查 APIResponse
taobao.wireless.content.check

无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70439.html" target="blank">阿里云内容安全</a>
*/
type TaobaoWirelessContentCheckAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWirelessContentCheckResponse `json:"wireless_content_check_response,omitempty"` 
    TaobaoWirelessContentCheckResponse
}

/* model for simplify = false
type TaobaoWirelessContentCheckResponse struct {

    // 检查结果
    
    CheckResults  struct {
        Checkpoints  []Checkpoints `json:"checkpoints,omitempty"`
    } `json:"check_results,omitempty"`
    

    // 综合结果建议。建议用户执行的操作，取值范围： pass：文本正常； review：需要人工审核； block：文本违规，可以直接删除或者做限制处理
    
    Suggestion   string `json:"suggestion,omitempty"`
    

}
*/

type TaobaoWirelessContentCheckResponse struct {

    // 检查结果
    CheckResults   []Checkpoints `json:"check_results,omitempty"`

    // 综合结果建议。建议用户执行的操作，取值范围： pass：文本正常； review：需要人工审核； block：文本违规，可以直接删除或者做限制处理
    Suggestion   string `json:"suggestion,omitempty"`

}
