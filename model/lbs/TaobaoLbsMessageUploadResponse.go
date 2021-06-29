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
    TaobaoLbsMessageUploadResponse
}

type TaobaoLbsMessageUploadResponse struct {
    XMLName xml.Name `xml:"lbs_message_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // subCode
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // subMsg
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
