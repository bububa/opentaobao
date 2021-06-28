package lbs

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
lbs数据采集 APIResponse
taobao.lbs.message.upload

lbs数据采集
*/
type TaobaoLbsMessageUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"lbs_message_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   bool `json:"result,omitempty" xml:"