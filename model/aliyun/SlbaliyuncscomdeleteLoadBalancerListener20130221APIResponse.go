package aliyun

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponseModel).Reset()
}

// SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponseModel is 删除 slb listener 成功返回结果
type SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponseModel struct {
	XMLName xml.Name `xml:"slb_aliyuncs_com_DeleteLoadBalancerListener_2013-02-21_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// request id
	Requestid string `json:"requestid,omitempty" xml:"requestid,omitempty"`
}

// Reset 清空结构体
func (m *SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponseModel) Reset() {
	m.RequestId = ""
	m.Requestid = ""
}

var poolSlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse = sync.Pool{
	New: func() any {
		return new(SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse)
	},
}

// GetSlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse 从 sync.Pool 获取 SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse
func GetSlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse() *SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse {
	return poolSlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse.Get().(*SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse)
}

// ReleaseSlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse 将 SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse 保存到 sync.Pool
func ReleaseSlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse(v *SlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse) {
	v.Reset()
	poolSlbAliyuncsComDeleteLoadBalancerListener20130221APIResponse.Put(v)
}
