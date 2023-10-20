package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse 物流工单履约状态更新 API返回值
// tmall.servicecenter.workcard.logisticsorder.update
//
// 天猫寄送类服务对接外部物流服务商回传物流状态信息
type TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardLogisticsorderUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardLogisticsorderUpdateAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardLogisticsorderUpdateAPIResponseModel is 物流工单履约状态更新 成功返回结果
type TmallServicecenterWorkcardLogisticsorderUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_logisticsorder_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *FulfilplatformResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardLogisticsorderUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardLogisticsorderUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse)
	},
}

// GetTmallServicecenterWorkcardLogisticsorderUpdateAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse
func GetTmallServicecenterWorkcardLogisticsorderUpdateAPIResponse() *TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse {
	return poolTmallServicecenterWorkcardLogisticsorderUpdateAPIResponse.Get().(*TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse)
}

// ReleaseTmallServicecenterWorkcardLogisticsorderUpdateAPIResponse 将 TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardLogisticsorderUpdateAPIResponse(v *TmallServicecenterWorkcardLogisticsorderUpdateAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardLogisticsorderUpdateAPIResponse.Put(v)
}
