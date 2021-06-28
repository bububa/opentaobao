package jms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询ONS分组 APIResponse
taobao.jushita.jms.group.get

查询当前appkey在ONS中已有的分组
*/
type TaobaoJushitaJmsGroupGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jushita_jms_group_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"