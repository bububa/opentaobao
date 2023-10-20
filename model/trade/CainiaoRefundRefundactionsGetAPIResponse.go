package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoRefundRefundactionsGetAPIResponse 判断该订单能执行的逆向操作集合列表 API返回值
// cainiao.refund.refundactions.get
//
// 判断该订单能执行的逆向操作集合列表
type CainiaoRefundRefundactionsGetAPIResponse struct {
	model.CommonResponse
	CainiaoRefundRefundactionsGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoRefundRefundactionsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoRefundRefundactionsGetAPIResponseModel).Reset()
}

// CainiaoRefundRefundactionsGetAPIResponseModel is 判断该订单能执行的逆向操作集合列表 成功返回结果
type CainiaoRefundRefundactionsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_refund_refundactions_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *CainiaoRefundRefundactionsGetBizResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoRefundRefundactionsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoRefundRefundactionsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoRefundRefundactionsGetAPIResponse)
	},
}

// GetCainiaoRefundRefundactionsGetAPIResponse 从 sync.Pool 获取 CainiaoRefundRefundactionsGetAPIResponse
func GetCainiaoRefundRefundactionsGetAPIResponse() *CainiaoRefundRefundactionsGetAPIResponse {
	return poolCainiaoRefundRefundactionsGetAPIResponse.Get().(*CainiaoRefundRefundactionsGetAPIResponse)
}

// ReleaseCainiaoRefundRefundactionsGetAPIResponse 将 CainiaoRefundRefundactionsGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoRefundRefundactionsGetAPIResponse(v *CainiaoRefundRefundactionsGetAPIResponse) {
	v.Reset()
	poolCainiaoRefundRefundactionsGetAPIResponse.Put(v)
}
