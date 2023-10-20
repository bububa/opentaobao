package degoperation

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationCheckAddrStatusAPIResponse 地址 API返回值
// taobao.degoperation.check.addr.status
//
// 激励
type TaobaoDegoperationCheckAddrStatusAPIResponse struct {
	model.CommonResponse
	TaobaoDegoperationCheckAddrStatusAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDegoperationCheckAddrStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDegoperationCheckAddrStatusAPIResponseModel).Reset()
}

// TaobaoDegoperationCheckAddrStatusAPIResponseModel is 地址 成功返回结果
type TaobaoDegoperationCheckAddrStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"degoperation_check_addr_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDegoperationCheckAddrStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoDegoperationCheckAddrStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDegoperationCheckAddrStatusAPIResponse)
	},
}

// GetTaobaoDegoperationCheckAddrStatusAPIResponse 从 sync.Pool 获取 TaobaoDegoperationCheckAddrStatusAPIResponse
func GetTaobaoDegoperationCheckAddrStatusAPIResponse() *TaobaoDegoperationCheckAddrStatusAPIResponse {
	return poolTaobaoDegoperationCheckAddrStatusAPIResponse.Get().(*TaobaoDegoperationCheckAddrStatusAPIResponse)
}

// ReleaseTaobaoDegoperationCheckAddrStatusAPIResponse 将 TaobaoDegoperationCheckAddrStatusAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDegoperationCheckAddrStatusAPIResponse(v *TaobaoDegoperationCheckAddrStatusAPIResponse) {
	v.Reset()
	poolTaobaoDegoperationCheckAddrStatusAPIResponse.Put(v)
}
