package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除 slb listener APIResponse
slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21

delete_vip
*/
type SlbAliyuncsComDeleteLoadBalancerListener2013-02-21APIResponse struct {
    model.CommonResponse
    // Response *SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Response `json:"slb_aliyuncs_com_DeleteLoadBalancerListener_2013-02-21_response,omitempty"` 
    SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Response
}

/* model for simplify = false
type SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Response struct {

    // request id
    
    Requestid   string `json:"requestid,omitempty"`
    

}
*/

type SlbAliyuncsComDeleteLoadBalancerListener2013-02-21Response struct {

    // request id
    Requestid   string `json:"requestid,omitempty"`

}
