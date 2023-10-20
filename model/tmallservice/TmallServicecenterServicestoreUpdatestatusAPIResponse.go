package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreUpdatestatusAPIResponse 网点/门店状态修改 API返回值
// tmall.servicecenter.servicestore.updatestatus
//
// 修改网点/门店状态
type TmallServicecenterServicestoreUpdatestatusAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreUpdatestatusAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreUpdatestatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicestoreUpdatestatusAPIResponseModel).Reset()
}

// TmallServicecenterServicestoreUpdatestatusAPIResponseModel is 网点/门店状态修改 成功返回结果
type TmallServicecenterServicestoreUpdatestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_updatestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 方法调用结果
	Result *TmallServicecenterServicestoreUpdatestatusResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreUpdatestatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicestoreUpdatestatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreUpdatestatusAPIResponse)
	},
}

// GetTmallServicecenterServicestoreUpdatestatusAPIResponse 从 sync.Pool 获取 TmallServicecenterServicestoreUpdatestatusAPIResponse
func GetTmallServicecenterServicestoreUpdatestatusAPIResponse() *TmallServicecenterServicestoreUpdatestatusAPIResponse {
	return poolTmallServicecenterServicestoreUpdatestatusAPIResponse.Get().(*TmallServicecenterServicestoreUpdatestatusAPIResponse)
}

// ReleaseTmallServicecenterServicestoreUpdatestatusAPIResponse 将 TmallServicecenterServicestoreUpdatestatusAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicestoreUpdatestatusAPIResponse(v *TmallServicecenterServicestoreUpdatestatusAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicestoreUpdatestatusAPIResponse.Put(v)
}
