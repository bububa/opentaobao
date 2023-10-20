package axintrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripAxinTransFundUpdateAPIResponse 修改资金单接口 API返回值
// taobao.alitrip.axin.trans.fund.update
//
// 阿信供销平台-修改资金单接口
type TaobaoAlitripAxinTransFundUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripAxinTransFundUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransFundUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripAxinTransFundUpdateAPIResponseModel).Reset()
}

// TaobaoAlitripAxinTransFundUpdateAPIResponseModel is 修改资金单接口 成功返回结果
type TaobaoAlitripAxinTransFundUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_axin_trans_fund_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoAlitripAxinTransFundUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripAxinTransFundUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAlitripAxinTransFundUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripAxinTransFundUpdateAPIResponse)
	},
}

// GetTaobaoAlitripAxinTransFundUpdateAPIResponse 从 sync.Pool 获取 TaobaoAlitripAxinTransFundUpdateAPIResponse
func GetTaobaoAlitripAxinTransFundUpdateAPIResponse() *TaobaoAlitripAxinTransFundUpdateAPIResponse {
	return poolTaobaoAlitripAxinTransFundUpdateAPIResponse.Get().(*TaobaoAlitripAxinTransFundUpdateAPIResponse)
}

// ReleaseTaobaoAlitripAxinTransFundUpdateAPIResponse 将 TaobaoAlitripAxinTransFundUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripAxinTransFundUpdateAPIResponse(v *TaobaoAlitripAxinTransFundUpdateAPIResponse) {
	v.Reset()
	poolTaobaoAlitripAxinTransFundUpdateAPIResponse.Put(v)
}
