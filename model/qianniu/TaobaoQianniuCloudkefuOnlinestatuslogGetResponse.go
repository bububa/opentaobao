package qianniu

import (
    "encoding/xml"

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
	RequestId     string         `json:"request_id,omitempty" xml:"qianniu_cloudkefu_onlinestatuslog_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // module
    
    RecordList   []RecordList `json:"record_list,omitempty" xml:"