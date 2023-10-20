package fans

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallFansCashpoolCheckpayAPIResponse 检查资金池付款状态 API返回值
// tmall.fans.cashpool.checkpay
//
// 检查资金池付款状态
type TmallFansCashpoolCheckpayAPIResponse struct {
	model.CommonResponse
	TmallFansCashpoolCheckpayAPIResponseModel
}

// Reset 清空结构体
func (m *TmallFansCashpoolCheckpayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallFansCashpoolCheckpayAPIResponseModel).Reset()
}

// TmallFansCashpoolCheckpayAPIResponseModel is 检查资金池付款状态 成功返回结果
type TmallFansCashpoolCheckpayAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fans_cashpool_checkpay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	FansResult *FansResult `json:"fans_result,omitempty" xml:"fans_result,omitempty"`
}

// Reset 清空结构体
func (m *TmallFansCashpoolCheckpayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FansResult = nil
}

var poolTmallFansCashpoolCheckpayAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallFansCashpoolCheckpayAPIResponse)
	},
}

// GetTmallFansCashpoolCheckpayAPIResponse 从 sync.Pool 获取 TmallFansCashpoolCheckpayAPIResponse
func GetTmallFansCashpoolCheckpayAPIResponse() *TmallFansCashpoolCheckpayAPIResponse {
	return poolTmallFansCashpoolCheckpayAPIResponse.Get().(*TmallFansCashpoolCheckpayAPIResponse)
}

// ReleaseTmallFansCashpoolCheckpayAPIResponse 将 TmallFansCashpoolCheckpayAPIResponse 保存到 sync.Pool
func ReleaseTmallFansCashpoolCheckpayAPIResponse(v *TmallFansCashpoolCheckpayAPIResponse) {
	v.Reset()
	poolTmallFansCashpoolCheckpayAPIResponse.Put(v)
}
