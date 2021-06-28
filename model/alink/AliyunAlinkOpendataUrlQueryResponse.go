package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
开放数据授权访问URL查询 APIResponse
aliyun.alink.opendata.url.query

厂商数据授权访问URL查询
*/
type AliyunAlinkOpendataUrlQueryAPIResponse struct {
    model.CommonResponse
    // Response *AliyunAlinkOpendataUrlQueryResponse `json:"aliyun_alink_opendata_url_query_response,omitempty"` 
    AliyunAlinkOpendataUrlQueryResponse
}

/* model for simplify = false
type AliyunAlinkOpendataUrlQueryResponse struct {

    // 调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 状态
    
    Status   int64 `json:"status,omitempty"`
    

    // 接口描述
    
    Message   string `json:"message,omitempty"`
    

    // 授权url
    
    Module  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"module,omitempty"`
    

}
*/

type AliyunAlinkOpendataUrlQueryResponse struct {

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 状态
    Status   int64 `json:"status,omitempty"`

    // 接口描述
    Message   string `json:"message,omitempty"`

    // 授权url
    Module   []string `json:"module,omitempty"`

}
