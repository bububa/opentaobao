package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkSpreadGetAPIResponse 淘宝客-公用-长链转短链 API返回值
// taobao.tbk.spread.get
//
// 输入一个原始的链接，转换得到指定的传播方式，如二维码，淘口令，短连接；
// 现阶段只支持短连接。
type TaobaoTbkSpreadGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkSpreadGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkSpreadGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkSpreadGetAPIResponseModel).Reset()
}

// TaobaoTbkSpreadGetAPIResponseModel is 淘宝客-公用-长链转短链 成功返回结果
type TaobaoTbkSpreadGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_spread_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 传播形式对象列表
	Results []TbkSpread `json:"results,omitempty" xml:"results>tbk_spread,omitempty"`
	// totalResults
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkSpreadGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.TotalResults = 0
}

var poolTaobaoTbkSpreadGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkSpreadGetAPIResponse)
	},
}

// GetTaobaoTbkSpreadGetAPIResponse 从 sync.Pool 获取 TaobaoTbkSpreadGetAPIResponse
func GetTaobaoTbkSpreadGetAPIResponse() *TaobaoTbkSpreadGetAPIResponse {
	return poolTaobaoTbkSpreadGetAPIResponse.Get().(*TaobaoTbkSpreadGetAPIResponse)
}

// ReleaseTaobaoTbkSpreadGetAPIResponse 将 TaobaoTbkSpreadGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkSpreadGetAPIResponse(v *TaobaoTbkSpreadGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkSpreadGetAPIResponse.Put(v)
}
