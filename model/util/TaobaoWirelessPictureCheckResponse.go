package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
无线开放图片内容安全检查 APIResponse
taobao.wireless.picture.check

无线开放内容检查，提供涉黄暴力政治图片检查。更详情介绍见 <a href="https://help.aliyun.com/document_detail/70292.html" target="blank">阿里云内容安全</a>
此API会进行两个场景审核，平均RT为1s。
*/
type TaobaoWirelessPictureCheckAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wireless_picture_check_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 检查结果
    
    CheckResults   []Checkpoints `json:"check_results,omitempty" xml:"