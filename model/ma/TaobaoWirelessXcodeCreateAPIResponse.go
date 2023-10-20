package ma

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWirelessXcodeCreateAPIResponse 创建二维码/短连接 API返回值
// taobao.wireless.xcode.create
//
// 创建码平台的普通二维码或者长连接转短连接服务
type TaobaoWirelessXcodeCreateAPIResponse struct {
	model.CommonResponse
	TaobaoWirelessXcodeCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWirelessXcodeCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWirelessXcodeCreateAPIResponseModel).Reset()
}

// TaobaoWirelessXcodeCreateAPIResponseModel is 创建二维码/短连接 成功返回结果
type TaobaoWirelessXcodeCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"wireless_xcode_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建二维码/短连接 返回信息
	Xcode *XCodeTo `json:"xcode,omitempty" xml:"xcode,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWirelessXcodeCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Xcode = nil
}

var poolTaobaoWirelessXcodeCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWirelessXcodeCreateAPIResponse)
	},
}

// GetTaobaoWirelessXcodeCreateAPIResponse 从 sync.Pool 获取 TaobaoWirelessXcodeCreateAPIResponse
func GetTaobaoWirelessXcodeCreateAPIResponse() *TaobaoWirelessXcodeCreateAPIResponse {
	return poolTaobaoWirelessXcodeCreateAPIResponse.Get().(*TaobaoWirelessXcodeCreateAPIResponse)
}

// ReleaseTaobaoWirelessXcodeCreateAPIResponse 将 TaobaoWirelessXcodeCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWirelessXcodeCreateAPIResponse(v *TaobaoWirelessXcodeCreateAPIResponse) {
	v.Reset()
	poolTaobaoWirelessXcodeCreateAPIResponse.Put(v)
}
