package aliyun

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// SlbAliyuncsComDescribeBackendServers20130221APIResponse 后端服务器健康检查 API返回值
// slb.aliyuncs.com.DescribeBackendServers.2013-02-21
//
// 后端服务器健康检查，对SLB实例的后端服务器进行健康检查，返回后端服务器的健康状况。后端服务器的健康状况为normal,abnormal和Unavailable三种。其中Unavailable表示这个SLB实例没有配置健康检查，无法获取后端服务器的健康状况。如果没有传入ListenerPort，则表示对这个SLB实例下的所有Listener后端服务器进行健康检查。
type SlbAliyuncsComDescribeBackendServers20130221APIResponse struct {
	model.CommonResponse
	SlbAliyuncsComDescribeBackendServers20130221APIResponseModel
}

// Reset 清空结构体
func (m *SlbAliyuncsComDescribeBackendServers20130221APIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.SlbAliyuncsComDescribeBackendServers20130221APIResponseModel).Reset()
}

// SlbAliyuncsComDescribeBackendServers20130221APIResponseModel is 后端服务器健康检查 成功返回结果
type SlbAliyuncsComDescribeBackendServers20130221APIResponseModel struct {
	XMLName xml.Name `xml:"slb_aliyuncs_com_DescribeBackendServers_2013-02-21_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Listeners
	Listeners []Listener `json:"listeners,omitempty" xml:"listeners>listener,omitempty"`
	// request id
	Requestid string `json:"requestid,omitempty" xml:"requestid,omitempty"`
	// LoadBalancerId
	Loadbalancerid string `json:"loadbalancerid,omitempty" xml:"loadbalancerid,omitempty"`
}

// Reset 清空结构体
func (m *SlbAliyuncsComDescribeBackendServers20130221APIResponseModel) Reset() {
	m.RequestId = ""
	m.Listeners = m.Listeners[:0]
	m.Requestid = ""
	m.Loadbalancerid = ""
}

var poolSlbAliyuncsComDescribeBackendServers20130221APIResponse = sync.Pool{
	New: func() any {
		return new(SlbAliyuncsComDescribeBackendServers20130221APIResponse)
	},
}

// GetSlbAliyuncsComDescribeBackendServers20130221APIResponse 从 sync.Pool 获取 SlbAliyuncsComDescribeBackendServers20130221APIResponse
func GetSlbAliyuncsComDescribeBackendServers20130221APIResponse() *SlbAliyuncsComDescribeBackendServers20130221APIResponse {
	return poolSlbAliyuncsComDescribeBackendServers20130221APIResponse.Get().(*SlbAliyuncsComDescribeBackendServers20130221APIResponse)
}

// ReleaseSlbAliyuncsComDescribeBackendServers20130221APIResponse 将 SlbAliyuncsComDescribeBackendServers20130221APIResponse 保存到 sync.Pool
func ReleaseSlbAliyuncsComDescribeBackendServers20130221APIResponse(v *SlbAliyuncsComDescribeBackendServers20130221APIResponse) {
	v.Reset()
	poolSlbAliyuncsComDescribeBackendServers20130221APIResponse.Put(v)
}
