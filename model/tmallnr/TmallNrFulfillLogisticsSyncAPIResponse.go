package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillLogisticsSyncAPIResponse 同城配物流信息回传 API返回值
// tmall.nr.fulfill.logistics.sync
//
// 同城配业务物流信息回传，通过接口将物流信息同步给天猫
type TmallNrFulfillLogisticsSyncAPIResponse struct {
	model.CommonResponse
	TmallNrFulfillLogisticsSyncAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrFulfillLogisticsSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrFulfillLogisticsSyncAPIResponseModel).Reset()
}

// TmallNrFulfillLogisticsSyncAPIResponseModel is 同城配物流信息回传 成功返回结果
type TmallNrFulfillLogisticsSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_fulfill_logistics_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrFulfillLogisticsSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrFulfillLogisticsSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrFulfillLogisticsSyncAPIResponse)
	},
}

// GetTmallNrFulfillLogisticsSyncAPIResponse 从 sync.Pool 获取 TmallNrFulfillLogisticsSyncAPIResponse
func GetTmallNrFulfillLogisticsSyncAPIResponse() *TmallNrFulfillLogisticsSyncAPIResponse {
	return poolTmallNrFulfillLogisticsSyncAPIResponse.Get().(*TmallNrFulfillLogisticsSyncAPIResponse)
}

// ReleaseTmallNrFulfillLogisticsSyncAPIResponse 将 TmallNrFulfillLogisticsSyncAPIResponse 保存到 sync.Pool
func ReleaseTmallNrFulfillLogisticsSyncAPIResponse(v *TmallNrFulfillLogisticsSyncAPIResponse) {
	v.Reset()
	poolTmallNrFulfillLogisticsSyncAPIResponse.Put(v)
}
