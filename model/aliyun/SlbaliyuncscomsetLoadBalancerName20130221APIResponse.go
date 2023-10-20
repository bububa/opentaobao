package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// SlbAliyuncsComSetLoadBalancerName20130221APIResponse 配置LoadBalancer的别名。 API返回值
// slb.aliyuncs.com.SetLoadBalancerName.2013-02-21
//
// 配置LoadBalancer的别名。
type SlbAliyuncsComSetLoadBalancerName20130221APIResponse struct {
	model.CommonResponse
	SlbAliyuncsComSetLoadBalancerName20130221APIResponseModel
}

// Reset 清空结构体
func (m *SlbAliyuncsComSetLoadBalancerName20130221APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.SlbAliyuncsComSetLoadBalancerName20130221APIResponseModel).Reset()
}

// SlbAliyuncsComSetLoadBalancerName20130221APIResponseModel is 配置LoadBalancer的别名。 成功返回结果
type SlbAliyuncsComSetLoadBalancerName20130221APIResponseModel struct {
	XMLName xml.Name `xml:"slb_aliyuncs_com_SetLoadBalancerName_2013-02-21_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// request id
	Requestid string `json:"requestid,omitempty" xml:"requestid,omitempty"`
}

// Reset 清空结构体
func (m *SlbAliyuncsComSetLoadBalancerName20130221APIResponseModel) Reset() {
	m.RequestId = ""
	m.Requestid = ""
}

var poolSlbAliyuncsComSetLoadBalancerName20130221APIResponse = sync.Pool{
	New: func() any {
		return new(SlbAliyuncsComSetLoadBalancerName20130221APIResponse)
	},
}

// GetSlbAliyuncsComSetLoadBalancerName20130221APIResponse 从 sync.Pool 获取 SlbAliyuncsComSetLoadBalancerName20130221APIResponse
func GetSlbAliyuncsComSetLoadBalancerName20130221APIResponse() *SlbAliyuncsComSetLoadBalancerName20130221APIResponse {
	return poolSlbAliyuncsComSetLoadBalancerName20130221APIResponse.Get().(*SlbAliyuncsComSetLoadBalancerName20130221APIResponse)
}

// ReleaseSlbAliyuncsComSetLoadBalancerName20130221APIResponse 将 SlbAliyuncsComSetLoadBalancerName20130221APIResponse 保存到 sync.Pool
func ReleaseSlbAliyuncsComSetLoadBalancerName20130221APIResponse(v *SlbAliyuncsComSetLoadBalancerName20130221APIResponse) {
	v.Reset()
	poolSlbAliyuncsComSetLoadBalancerName20130221APIResponse.Put(v)
}
