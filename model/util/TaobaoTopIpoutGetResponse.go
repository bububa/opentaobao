package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取开放平台出口IP段 APIResponse
taobao.top.ipout.get

获取开放平台出口IP段
*/
type TaobaoTopIpoutGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"top_ipout_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // TOP网关出口IP列表
    
    IpList   string `json:"ip_list,omitempty" xml:"