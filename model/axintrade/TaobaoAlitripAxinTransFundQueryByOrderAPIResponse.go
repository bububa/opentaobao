package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransFundQueryByOrderAPIResponse 通过外部订单ID查询所有资金单 API返回值
// taobao.alitrip.axin.trans.fund.query.by.order
//
// 阿信供销平台-通过外部订单ID查询所有资金单
type TaobaoAlitripAxinTransFundQueryByOrderAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransFundQueryByOrderAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransFundQueryByOrderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripAxinTransFundQueryByOrderAPIResponseModel).Reset()
}

// TaobaoAlitripAxinTransFundQueryByOrderAPIResponseModel is 通过外部订单ID查询所有资金单 成功返回结果
type TaobaoAlitripAxinTransFundQueryByOrderAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_fund_query_by_order_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripAxinTransFundQueryByOrderResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransFundQueryByOrderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripAxinTransFundQueryByOrderAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransFundQueryByOrderAPIResponse)
	},
}

// GetTaobaoAlitripAxinTransFundQueryByOrderAPIResponse 从 sync.Pool 获取 TaobaoAlitripAxinTransFundQueryByOrderAPIResponse
func GetTaobaoAlitripAxinTransFundQueryByOrderAPIResponse() *TaobaoAlitripAxinTransFundQueryByOrderAPIResponse {
	return poolTaobaoAlitripAxinTransFundQueryByOrderAPIResponse.Get().(*TaobaoAlitripAxinTransFundQueryByOrderAPIResponse)
}

// ReleaseTaobaoAlitripAxinTransFundQueryByOrderAPIResponse 将 TaobaoAlitripAxinTransFundQueryByOrderAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripAxinTransFundQueryByOrderAPIResponse(v *TaobaoAlitripAxinTransFundQueryByOrderAPIResponse) {
	v.Reset()
	poolTaobaoAlitripAxinTransFundQueryByOrderAPIResponse.Put(v)
}
