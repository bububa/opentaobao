package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenuidGetAPIResponse 获取授权账号对应的OpenUid API返回值
// taobao.openuid.get
//
// 获取授权账号对应的OpenUid
type TaobaoOpenuidGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenuidGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenuidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenuidGetAPIResponseModel).Reset()
}

// TaobaoOpenuidGetAPIResponseModel is 获取授权账号对应的OpenUid 成功返回结果
type TaobaoOpenuidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openuid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// OpenUID
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenuidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OpenUid = ""
}

var poolTaobaoOpenuidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenuidGetAPIResponse)
	},
}

// GetTaobaoOpenuidGetAPIResponse 从 sync.Pool 获取 TaobaoOpenuidGetAPIResponse
func GetTaobaoOpenuidGetAPIResponse() *TaobaoOpenuidGetAPIResponse {
	return poolTaobaoOpenuidGetAPIResponse.Get().(*TaobaoOpenuidGetAPIResponse)
}

// ReleaseTaobaoOpenuidGetAPIResponse 将 TaobaoOpenuidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenuidGetAPIResponse(v *TaobaoOpenuidGetAPIResponse) {
	v.Reset()
	poolTaobaoOpenuidGetAPIResponse.Put(v)
}
