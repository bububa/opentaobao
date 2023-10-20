package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoRefundRefundactionsJudgementAPIResponse 判断当前用户是否能对订单执行一些逆向操作，比如退货操作 API返回值
// cainiao.refund.refundactions.judgement
//
// 判断当前用户是否能对订单执行一些逆向操作，比如退货操作
type CainiaoRefundRefundactionsJudgementAPIResponse struct {
	model.CommonResponse
	CainiaoRefundRefundactionsJudgementAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoRefundRefundactionsJudgementAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoRefundRefundactionsJudgementAPIResponseModel).Reset()
}

// CainiaoRefundRefundactionsJudgementAPIResponseModel is 判断当前用户是否能对订单执行一些逆向操作，比如退货操作 成功返回结果
type CainiaoRefundRefundactionsJudgementAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_refund_refundactions_judgement_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *CainiaoRefundRefundactionsJudgementBizResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoRefundRefundactionsJudgementAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolCainiaoRefundRefundactionsJudgementAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoRefundRefundactionsJudgementAPIResponse)
	},
}

// GetCainiaoRefundRefundactionsJudgementAPIResponse 从 sync.Pool 获取 CainiaoRefundRefundactionsJudgementAPIResponse
func GetCainiaoRefundRefundactionsJudgementAPIResponse() *CainiaoRefundRefundactionsJudgementAPIResponse {
	return poolCainiaoRefundRefundactionsJudgementAPIResponse.Get().(*CainiaoRefundRefundactionsJudgementAPIResponse)
}

// ReleaseCainiaoRefundRefundactionsJudgementAPIResponse 将 CainiaoRefundRefundactionsJudgementAPIResponse 保存到 sync.Pool
func ReleaseCainiaoRefundRefundactionsJudgementAPIResponse(v *CainiaoRefundRefundactionsJudgementAPIResponse) {
	v.Reset()
	poolCainiaoRefundRefundactionsJudgementAPIResponse.Put(v)
}
