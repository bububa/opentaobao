package aliyun

import (
    "github.com/bububa/opentaobao/model"
)

/* 
配置LoadBalancer的别名。 APIResponse
slb.aliyuncs.com.SetLoadBalancerName.2013-02-21

配置LoadBalancer的别名。
*/
type SlbAliyuncsComSetLoadBalancerName2013-02-21APIResponse struct {
    model.CommonResponse
    Response *SlbAliyuncsComSetLoadBalancerName2013-02-21Response `json:"slb_aliyuncs_com_SetLoadBalancerName_2013-02-21_response,omitempty"`
}

type SlbAliyuncsComSetLoadBalancerName2013-02-21Response struct {

    // request id
    Requestid   string `json:"requestid,omitempty"`

}
