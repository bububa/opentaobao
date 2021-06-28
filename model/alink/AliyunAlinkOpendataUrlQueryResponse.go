package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
开放数据授权访问URL查询 APIResponse
aliyun.alink.opendata.url.query

厂商数据授权访问URL查询
*/
type AliyunAlinkOpendataUrlQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"aliyun_alink_opendata_url_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"