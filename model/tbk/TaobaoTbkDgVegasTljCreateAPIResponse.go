package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasTljCreateAPIResponse 淘宝客-推广者-淘礼金创建 API返回值
// taobao.tbk.dg.vegas.tlj.create
//
// 创建淘礼金
type TaobaoTbkDgVegasTljCreateAPIResponse struct {
	model.CommonResponse
	TaobaoTbkDgVegasTljCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkDgVegasTljCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkDgVegasTljCreateAPIResponseModel).Reset()
}

// TaobaoTbkDgVegasTljCreateAPIResponseModel is 淘宝客-推广者-淘礼金创建 成功返回结果
type TaobaoTbkDgVegasTljCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_dg_vegas_tlj_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTbkDgVegasTljCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkDgVegasTljCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTbkDgVegasTljCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkDgVegasTljCreateAPIResponse)
	},
}

// GetTaobaoTbkDgVegasTljCreateAPIResponse 从 sync.Pool 获取 TaobaoTbkDgVegasTljCreateAPIResponse
func GetTaobaoTbkDgVegasTljCreateAPIResponse() *TaobaoTbkDgVegasTljCreateAPIResponse {
	return poolTaobaoTbkDgVegasTljCreateAPIResponse.Get().(*TaobaoTbkDgVegasTljCreateAPIResponse)
}

// ReleaseTaobaoTbkDgVegasTljCreateAPIResponse 将 TaobaoTbkDgVegasTljCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkDgVegasTljCreateAPIResponse(v *TaobaoTbkDgVegasTljCreateAPIResponse) {
	v.Reset()
	poolTaobaoTbkDgVegasTljCreateAPIResponse.Put(v)
}
