package tanx

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxDealGetAPIResponse 对外部dsp提供交易id查询接口 API返回值
// taobao.tanx.deal.get
//
// 对外部dsp提供交易id查询接口
type TaobaoTanxDealGetAPIResponse struct {
	model.CommonResponse
	TaobaoTanxDealGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTanxDealGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTanxDealGetAPIResponseModel).Reset()
}

// TaobaoTanxDealGetAPIResponseModel is 对外部dsp提供交易id查询接口 成功返回结果
type TaobaoTanxDealGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_deal_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果信息
	Messag string `json:"messag,omitempty" xml:"messag,omitempty"`
	// 结果代码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 查询结果
	Result *DealInfoDto `json:"result,omitempty" xml:"result,omitempty"`
	// 查询结果
	Sucess bool `json:"sucess,omitempty" xml:"sucess,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTanxDealGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Messag = ""
	m.Code = 0
	m.Result = nil
	m.Sucess = false
}

var poolTaobaoTanxDealGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTanxDealGetAPIResponse)
	},
}

// GetTaobaoTanxDealGetAPIResponse 从 sync.Pool 获取 TaobaoTanxDealGetAPIResponse
func GetTaobaoTanxDealGetAPIResponse() *TaobaoTanxDealGetAPIResponse {
	return poolTaobaoTanxDealGetAPIResponse.Get().(*TaobaoTanxDealGetAPIResponse)
}

// ReleaseTaobaoTanxDealGetAPIResponse 将 TaobaoTanxDealGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTanxDealGetAPIResponse(v *TaobaoTanxDealGetAPIResponse) {
	v.Reset()
	poolTaobaoTanxDealGetAPIResponse.Put(v)
}
