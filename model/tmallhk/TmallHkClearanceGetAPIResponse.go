package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallHkClearanceGetAPIResponse 天猫国际-清关材料查询 API返回值
// tmall.hk.clearance.get
//
// 提供订单收货人身份信息查询功能。
type TmallHkClearanceGetAPIResponse struct {
	model.CommonResponse
	TmallHkClearanceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallHkClearanceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallHkClearanceGetAPIResponseModel).Reset()
}

// TmallHkClearanceGetAPIResponseModel is 天猫国际-清关材料查询 成功返回结果
type TmallHkClearanceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_hk_clearance_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果对象
	Result *CertifyQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallHkClearanceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallHkClearanceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallHkClearanceGetAPIResponse)
	},
}

// GetTmallHkClearanceGetAPIResponse 从 sync.Pool 获取 TmallHkClearanceGetAPIResponse
func GetTmallHkClearanceGetAPIResponse() *TmallHkClearanceGetAPIResponse {
	return poolTmallHkClearanceGetAPIResponse.Get().(*TmallHkClearanceGetAPIResponse)
}

// ReleaseTmallHkClearanceGetAPIResponse 将 TmallHkClearanceGetAPIResponse 保存到 sync.Pool
func ReleaseTmallHkClearanceGetAPIResponse(v *TmallHkClearanceGetAPIResponse) {
	v.Reset()
	poolTmallHkClearanceGetAPIResponse.Put(v)
}
