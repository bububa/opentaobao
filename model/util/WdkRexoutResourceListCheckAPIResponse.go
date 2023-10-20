package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkRexoutResourceListCheckAPIResponse ReX应用中心资源更新检测-外部 API返回值
// wdk.rexout.resource.list.check
//
// ReX应用中心资源更新检测-外部
type WdkRexoutResourceListCheckAPIResponse struct {
	model.CommonResponse
	WdkRexoutResourceListCheckAPIResponseModel
}

// Reset 清空结构体
func (m *WdkRexoutResourceListCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkRexoutResourceListCheckAPIResponseModel).Reset()
}

// WdkRexoutResourceListCheckAPIResponseModel is ReX应用中心资源更新检测-外部 成功返回结果
type WdkRexoutResourceListCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_rexout_resource_list_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 资源列表
	Data []Dto `json:"data,omitempty" xml:"data>dto,omitempty"`
	// 是否成功
	Succeed string `json:"succeed,omitempty" xml:"succeed,omitempty"`
	// 错误内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}

// Reset 清空结构体
func (m *WdkRexoutResourceListCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.Succeed = ""
	m.MsgInfo = ""
	m.MsgCode = ""
}

var poolWdkRexoutResourceListCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkRexoutResourceListCheckAPIResponse)
	},
}

// GetWdkRexoutResourceListCheckAPIResponse 从 sync.Pool 获取 WdkRexoutResourceListCheckAPIResponse
func GetWdkRexoutResourceListCheckAPIResponse() *WdkRexoutResourceListCheckAPIResponse {
	return poolWdkRexoutResourceListCheckAPIResponse.Get().(*WdkRexoutResourceListCheckAPIResponse)
}

// ReleaseWdkRexoutResourceListCheckAPIResponse 将 WdkRexoutResourceListCheckAPIResponse 保存到 sync.Pool
func ReleaseWdkRexoutResourceListCheckAPIResponse(v *WdkRexoutResourceListCheckAPIResponse) {
	v.Reset()
	poolWdkRexoutResourceListCheckAPIResponse.Put(v)
}
