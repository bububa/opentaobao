package aliyun

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse 删除 slb listener API返回值
// slb.aliyuncs.com.DeleteLoadBalancerListener.2013-02-21
//
// delete_vip
type SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse struct {
	model.CommonResponse
	SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponseModel
}

// SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponseModel is 删除 slb listener 成功返回结果
type SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponseModel struct {
	XMLName xml.Name `xml:"slb_aliyuncs_com_DeleteLoadBalancerListener_2013-02-21_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// request id
	Requestid string `json:"requestid,omitempty" xml:"requestid,omitempty"`
}
