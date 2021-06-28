package aliyun

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除 slb listener APIResponse
slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21

delete_vip
*/
type SlbAliyuncsComDeleteLoadBalancerListener2013-02-21APIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"slb_aliyuncs_com_DeleteLoadBalancerListener_2013-02-21_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // request id
    
    Requestid   string `json:"requestid,omitempty" xml:"