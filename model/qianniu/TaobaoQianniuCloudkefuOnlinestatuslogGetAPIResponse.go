package qianniu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse 查询客服在线状态 API返回值
// taobao.qianniu.cloudkefu.onlinestatuslog.get
//
// 按天查询客服账号的在线状态记录。如：登录，下线，挂起等
// 有别于taobao.qianniu.cloudkefu.statuslog.get接口，这个接口可以查询30天内的流水，不需要分页查询
type TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse struct {
	model.CommonResponse
	TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponseModel).Reset()
}

// TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponseModel is 查询客服在线状态 成功返回结果
type TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qianniu_cloudkefu_onlinestatuslog_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// module
	RecordList []RecordList `json:"record_list,omitempty" xml:"record_list>record_list,omitempty"`
	// cause
	Cause string `json:"cause,omitempty" xml:"cause,omitempty"`
	// errorMap
	ErrorMap string `json:"error_map,omitempty" xml:"error_map,omitempty"`
	// attachment
	Attachment string `json:"attachment,omitempty" xml:"attachment,omitempty"`
	// solution
	Solution string `json:"solution,omitempty" xml:"solution,omitempty"`
	// version
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RecordList = m.RecordList[:0]
	m.Cause = ""
	m.ErrorMap = ""
	m.Attachment = ""
	m.Solution = ""
	m.Version = 0
}

var poolTaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse)
	},
}

// GetTaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse 从 sync.Pool 获取 TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse
func GetTaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse() *TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse {
	return poolTaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse.Get().(*TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse)
}

// ReleaseTaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse 将 TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse(v *TaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse) {
	v.Reset()
	poolTaobaoQianniuCloudkefuOnlinestatuslogGetAPIResponse.Put(v)
}
