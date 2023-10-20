package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse tp更新物流进度信息 API返回值
// tmall.servicecenter.workcard.logisticsorder.tpcancel
//
// tp更新物流进度信息
type TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponseModel is tp更新物流进度信息 成功返回结果
type TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_logisticsorder_tpcancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse)
	},
}

// GetTmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse
func GetTmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse() *TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse {
	return poolTmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse.Get().(*TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse)
}

// ReleaseTmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse 将 TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse(v *TmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardLogisticsorderTpcancelAPIResponse.Put(v)
}
