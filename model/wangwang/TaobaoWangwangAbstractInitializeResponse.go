package wangwang

import (
    "github.com/bububa/opentaobao/model"
)

/* 
模糊查询服务初始化 APIResponse
taobao.wangwang.abstract.initialize

模糊查询服务初始化，只支持json返回
*/
type TaobaoWangwangAbstractInitializeAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWangwangAbstractInitializeResponse `json:"taobao_wangwang_abstract_initialize_response,omitempty"`
}

type TaobaoWangwangAbstractInitializeResponse struct {

    // 0或-1表示成功或失败
    RetCode   int64 `json:"ret_code,omitempty"`

    // 当ret_code=-1时这个变量才有
    ErrorMsg   string `json:"error_msg,omitempty"`

}
