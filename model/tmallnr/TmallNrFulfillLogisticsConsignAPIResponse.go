package tmallnr

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrFulfillLogisticsConsignAPIResponse 同城配门店备货通知 API返回值
// tmall.nr.fulfill.logistics.consign
//
// 同城配业务备货通知，商家告诉平台门店的货已经准备好，可以发货了；
type TmallNrFulfillLogisticsConsignAPIResponse struct {
	model.CommonResponse
	TmallNrFulfillLogisticsConsignAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrFulfillLogisticsConsignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrFulfillLogisticsConsignAPIResponseModel).Reset()
}

// TmallNrFulfillLogisticsConsignAPIResponseModel is 同城配门店备货通知 成功返回结果
type TmallNrFulfillLogisticsConsignAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_fulfill_logistics_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrFulfillLogisticsConsignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallNrFulfillLogisticsConsignAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrFulfillLogisticsConsignAPIResponse)
	},
}

// GetTmallNrFulfillLogisticsConsignAPIResponse 从 sync.Pool 获取 TmallNrFulfillLogisticsConsignAPIResponse
func GetTmallNrFulfillLogisticsConsignAPIResponse() *TmallNrFulfillLogisticsConsignAPIResponse {
	return poolTmallNrFulfillLogisticsConsignAPIResponse.Get().(*TmallNrFulfillLogisticsConsignAPIResponse)
}

// ReleaseTmallNrFulfillLogisticsConsignAPIResponse 将 TmallNrFulfillLogisticsConsignAPIResponse 保存到 sync.Pool
func ReleaseTmallNrFulfillLogisticsConsignAPIResponse(v *TmallNrFulfillLogisticsConsignAPIResponse) {
	v.Reset()
	poolTmallNrFulfillLogisticsConsignAPIResponse.Put(v)
}
