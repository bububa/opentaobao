package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardCallRecordAPIResponse 回访记录回传API API返回值
// tmall.servicecenter.workcard.call.record
//
// 客满回访信息登记
type TmallServicecenterWorkcardCallRecordAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardCallRecordAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardCallRecordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardCallRecordAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardCallRecordAPIResponseModel is 回访记录回传API 成功返回结果
type TmallServicecenterWorkcardCallRecordAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_call_record_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardCallRecordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardCallRecordAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardCallRecordAPIResponse)
	},
}

// GetTmallServicecenterWorkcardCallRecordAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardCallRecordAPIResponse
func GetTmallServicecenterWorkcardCallRecordAPIResponse() *TmallServicecenterWorkcardCallRecordAPIResponse {
	return poolTmallServicecenterWorkcardCallRecordAPIResponse.Get().(*TmallServicecenterWorkcardCallRecordAPIResponse)
}

// ReleaseTmallServicecenterWorkcardCallRecordAPIResponse 将 TmallServicecenterWorkcardCallRecordAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardCallRecordAPIResponse(v *TmallServicecenterWorkcardCallRecordAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardCallRecordAPIResponse.Put(v)
}
