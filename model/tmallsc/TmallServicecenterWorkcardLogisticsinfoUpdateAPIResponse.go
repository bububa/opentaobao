package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse 物流单信息回传接口 API返回值
// tmall.servicecenter.workcard.logisticsinfo.update
//
// 物流单信息回传接口
type TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponseModel is 物流单信息回传接口 成功返回结果
type TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_logisticsinfo_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.IsSuccess = false
}

var poolTmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse)
	},
}

// GetTmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse
func GetTmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse() *TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse {
	return poolTmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse.Get().(*TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse)
}

// ReleaseTmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse 将 TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse(v *TmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardLogisticsinfoUpdateAPIResponse.Put(v)
}
