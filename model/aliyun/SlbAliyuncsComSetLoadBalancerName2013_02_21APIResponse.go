package aliyun

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// SlbAliyuncsComSetLoadBalancerName2013_02_21APIResponse 配置LoadBalancer的别名。 API返回值
// slb.aliyuncs.com.SetLoadBalancerName.2013-02-21
//
// 配置LoadBalancer的别名。
type SlbAliyuncsComSetLoadBalancerName2013_02_21APIResponse struct {
	model.CommonResponse
	SlbAliyuncsComSetLoadBalancerName2013_02_21APIResponseModel
}

// SlbAliyuncsComSetLoadBalancerName2013_02_21APIResponseModel is 配置LoadBalancer的别名。 成功返回结果
type SlbAliyuncsComSetLoadBalancerName2013_02_21APIResponseModel struct {
	XMLName xml.Name `xml:"slb_aliyuncs_com_SetLoadBalancerName_2013-02-21_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// request id
	Requestid string `json:"requestid,omitempty" xml:"requestid,omitempty"`
}
