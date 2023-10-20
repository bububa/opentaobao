package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoAddAPIResponse 自促活动申请接口 API返回值
// taobao.xhotel.bnbpromo.add
//
// 自促活动申请接口
type TaobaoXhotelBnbpromoAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbpromoAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbpromoAddAPIResponseModel).Reset()
}

// TaobaoXhotelBnbpromoAddAPIResponseModel is 自促活动申请接口 成功返回结果
type TaobaoXhotelBnbpromoAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销添加返回对象
	Module *PromoCode `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = nil
}

var poolTaobaoXhotelBnbpromoAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbpromoAddAPIResponse)
	},
}

// GetTaobaoXhotelBnbpromoAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbpromoAddAPIResponse
func GetTaobaoXhotelBnbpromoAddAPIResponse() *TaobaoXhotelBnbpromoAddAPIResponse {
	return poolTaobaoXhotelBnbpromoAddAPIResponse.Get().(*TaobaoXhotelBnbpromoAddAPIResponse)
}

// ReleaseTaobaoXhotelBnbpromoAddAPIResponse 将 TaobaoXhotelBnbpromoAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbpromoAddAPIResponse(v *TaobaoXhotelBnbpromoAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbpromoAddAPIResponse.Put(v)
}
