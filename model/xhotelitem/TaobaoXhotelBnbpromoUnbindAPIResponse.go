package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoUnbindAPIResponse 自促活动解绑接口 API返回值
// taobao.xhotel.bnbpromo.unbind
//
// 自促活动解绑接口
type TaobaoXhotelBnbpromoUnbindAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbpromoUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbpromoUnbindAPIResponseModel).Reset()
}

// TaobaoXhotelBnbpromoUnbindAPIResponseModel is 自促活动解绑接口 成功返回结果
type TaobaoXhotelBnbpromoUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销解绑返回对象
	Module *PromoBindResult `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = nil
}

var poolTaobaoXhotelBnbpromoUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbpromoUnbindAPIResponse)
	},
}

// GetTaobaoXhotelBnbpromoUnbindAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbpromoUnbindAPIResponse
func GetTaobaoXhotelBnbpromoUnbindAPIResponse() *TaobaoXhotelBnbpromoUnbindAPIResponse {
	return poolTaobaoXhotelBnbpromoUnbindAPIResponse.Get().(*TaobaoXhotelBnbpromoUnbindAPIResponse)
}

// ReleaseTaobaoXhotelBnbpromoUnbindAPIResponse 将 TaobaoXhotelBnbpromoUnbindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbpromoUnbindAPIResponse(v *TaobaoXhotelBnbpromoUnbindAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbpromoUnbindAPIResponse.Put(v)
}
