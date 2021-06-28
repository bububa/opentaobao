package wangwang

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
模糊查询服务初始化 APIResponse
taobao.wangwang.abstract.initialize

模糊查询服务初始化，只支持json返回
*/
type TaobaoWangwangAbstractInitializeAPIResponse struct {
    model.CommonResponse
    TaobaoWangwangAbstractInitializeResponse
}

type TaobaoWangwangAbstractInitializeResponse struct {
    XMLName xml.Name `xml:"wangwang_abstract_initialize_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 0或-1表示成功或失败
    
    RetCode   int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`

    
    // 当ret_code=-1时这个变量才有
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
