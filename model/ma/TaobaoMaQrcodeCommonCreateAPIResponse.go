package ma

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMaQrcodeCommonCreateAPIResponse 创建码平台常用二维码 API返回值
// taobao.ma.qrcode.common.create
//
// 创建码平台对外提供的常用二维码接口，不适于码平台业务类型的码创建，如不支持包裹码、媒体码等，业务类型的码需要单独提供API。
type TaobaoMaQrcodeCommonCreateAPIResponse struct {
	model.CommonResponse
	TaobaoMaQrcodeCommonCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMaQrcodeCommonCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMaQrcodeCommonCreateAPIResponseModel).Reset()
}

// TaobaoMaQrcodeCommonCreateAPIResponseModel is 创建码平台常用二维码 成功返回结果
type TaobaoMaQrcodeCommonCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"ma_qrcode_common_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 二维码对像
	Modules []QrcodeDo `json:"modules,omitempty" xml:"modules>qrcode_do,omitempty"`
	// 执行是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMaQrcodeCommonCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Modules = m.Modules[:0]
	m.Suc = false
}

var poolTaobaoMaQrcodeCommonCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMaQrcodeCommonCreateAPIResponse)
	},
}

// GetTaobaoMaQrcodeCommonCreateAPIResponse 从 sync.Pool 获取 TaobaoMaQrcodeCommonCreateAPIResponse
func GetTaobaoMaQrcodeCommonCreateAPIResponse() *TaobaoMaQrcodeCommonCreateAPIResponse {
	return poolTaobaoMaQrcodeCommonCreateAPIResponse.Get().(*TaobaoMaQrcodeCommonCreateAPIResponse)
}

// ReleaseTaobaoMaQrcodeCommonCreateAPIResponse 将 TaobaoMaQrcodeCommonCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMaQrcodeCommonCreateAPIResponse(v *TaobaoMaQrcodeCommonCreateAPIResponse) {
	v.Reset()
	poolTaobaoMaQrcodeCommonCreateAPIResponse.Put(v)
}
