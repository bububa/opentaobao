package qianniu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询客服在线状态 APIResponse
taobao.qianniu.cloudkefu.onlinestatuslog.get

按天查询客服账号的在线状态记录。如：登录，下线，挂起等
有别于taobao.qianniu.cloudkefu.statuslog.get接口，这个接口可以查询30天内的流水，不需要分页查询
*/
type TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQianniuCloudkefuOnlinestatuslogGetResponse `json:"qianniu_cloudkefu_onlinestatuslog_get_response,omitempty"` 
    TaobaoQianniuCloudkefuOnlinestatuslogGetResponse
}

/* model for simplify = false
type TaobaoQianniuCloudkefuOnlinestatuslogGetResponse struct {

    // module
    
    RecordList  struct {
        RecordList  []RecordList `json:"record_list,omitempty"`
    } `json:"record_list,omitempty"`
    

    // cause
    
    Cause   string `json:"cause,omitempty"`
    

    // errorMap
    
    ErrorMap   string `json:"error_map,omitempty"`
    

    // attachment
    
    Attachment   string `json:"attachment,omitempty"`
    

    // solution
    
    Solution   string `json:"solution,omitempty"`
    

    // version
    
    Version   int64 `json:"version,omitempty"`
    

}
*/

type TaobaoQianniuCloudkefuOnlinestatuslogGetResponse struct {

    // module
    RecordList   []RecordList `json:"record_list,omitempty"`

    // cause
    Cause   string `json:"cause,omitempty"`

    // errorMap
    ErrorMap   string `json:"error_map,omitempty"`

    // attachment
    Attachment   string `json:"attachment,omitempty"`

    // solution
    Solution   string `json:"solution,omitempty"`

    // version
    Version   int64 `json:"version,omitempty"`

}
