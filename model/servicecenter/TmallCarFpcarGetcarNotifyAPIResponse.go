package servicecenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarFpcarGetcarNotifyAPIResponse 门店通知用户提车 API返回值
// tmall.car.fpcar.getcar.notify
//
// 提供给外部(大搜或其它合作方)的接口-门店通知用户提车
type TmallCarFpcarGetcarNotifyAPIResponse struct {
	model.CommonResponse
	TmallCarFpcarGetcarNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarFpcarGetcarNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarFpcarGetcarNotifyAPIResponseModel).Reset()
}

// TmallCarFpcarGetcarNotifyAPIResponseModel is 门店通知用户提车 成功返回结果
type TmallCarFpcarGetcarNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_fpcar_getcar_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据结果
	Object string `json:"object,omitempty" xml:"object,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Succes bool `json:"succes,omitempty" xml:"succes,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarFpcarGetcarNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Object = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Succes = false
}

var poolTmallCarFpcarGetcarNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarFpcarGetcarNotifyAPIResponse)
	},
}

// GetTmallCarFpcarGetcarNotifyAPIResponse 从 sync.Pool 获取 TmallCarFpcarGetcarNotifyAPIResponse
func GetTmallCarFpcarGetcarNotifyAPIResponse() *TmallCarFpcarGetcarNotifyAPIResponse {
	return poolTmallCarFpcarGetcarNotifyAPIResponse.Get().(*TmallCarFpcarGetcarNotifyAPIResponse)
}

// ReleaseTmallCarFpcarGetcarNotifyAPIResponse 将 TmallCarFpcarGetcarNotifyAPIResponse 保存到 sync.Pool
func ReleaseTmallCarFpcarGetcarNotifyAPIResponse(v *TmallCarFpcarGetcarNotifyAPIResponse) {
	v.Reset()
	poolTmallCarFpcarGetcarNotifyAPIResponse.Put(v)
}
