package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
无线开放内容检查 APIResponse
taobao.wireless.content.check

无线开放内容检查，提供涉黄暴力政治文本检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70439.html" target="blank">阿里云内容安全</a>
*/
type TaobaoWirelessContentCheckAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wireless_content_check_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 检查结果
    
    CheckResults   []Checkpoints `json:"check_results,omitempty" xml:"