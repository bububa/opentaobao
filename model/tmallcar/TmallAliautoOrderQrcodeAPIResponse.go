package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoOrderQrcodeAPIResponse 根据商品id列表获取可扫描下单二维码 API返回值
// tmall.aliauto.order.qrcode
//
// 根据商品id列表获取可扫描下单二维码
type TmallAliautoOrderQrcodeAPIResponse struct {
	model.CommonResponse
	TmallAliautoOrderQrcodeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallAliautoOrderQrcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallAliautoOrderQrcodeAPIResponseModel).Reset()
}

// TmallAliautoOrderQrcodeAPIResponseModel is 根据商品id列表获取可扫描下单二维码 成功返回结果
type TmallAliautoOrderQrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_aliauto_order_qrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的数据实体
	Data *ConfirmOrderQrCode4IsvDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TmallAliautoOrderQrcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTmallAliautoOrderQrcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallAliautoOrderQrcodeAPIResponse)
	},
}

// GetTmallAliautoOrderQrcodeAPIResponse 从 sync.Pool 获取 TmallAliautoOrderQrcodeAPIResponse
func GetTmallAliautoOrderQrcodeAPIResponse() *TmallAliautoOrderQrcodeAPIResponse {
	return poolTmallAliautoOrderQrcodeAPIResponse.Get().(*TmallAliautoOrderQrcodeAPIResponse)
}

// ReleaseTmallAliautoOrderQrcodeAPIResponse 将 TmallAliautoOrderQrcodeAPIResponse 保存到 sync.Pool
func ReleaseTmallAliautoOrderQrcodeAPIResponse(v *TmallAliautoOrderQrcodeAPIResponse) {
	v.Reset()
	poolTmallAliautoOrderQrcodeAPIResponse.Put(v)
}
