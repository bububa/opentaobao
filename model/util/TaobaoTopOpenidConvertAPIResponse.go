package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopOpenidConvertAPIResponse 混淆nick转openid API返回值
// taobao.top.openid.convert
//
// 混淆nick转openid，生成混淆nick必须与当前请求的isv匹配
type TaobaoTopOpenidConvertAPIResponse struct {
	model.CommonResponse
	TaobaoTopOpenidConvertAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopOpenidConvertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopOpenidConvertAPIResponseModel).Reset()
}

// TaobaoTopOpenidConvertAPIResponseModel is 混淆nick转openid 成功返回结果
type TaobaoTopOpenidConvertAPIResponseModel struct {
	XMLName xml.Name `xml:"top_openid_convert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// open_id
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopOpenidConvertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OpenId = ""
}

var poolTaobaoTopOpenidConvertAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopOpenidConvertAPIResponse)
	},
}

// GetTaobaoTopOpenidConvertAPIResponse 从 sync.Pool 获取 TaobaoTopOpenidConvertAPIResponse
func GetTaobaoTopOpenidConvertAPIResponse() *TaobaoTopOpenidConvertAPIResponse {
	return poolTaobaoTopOpenidConvertAPIResponse.Get().(*TaobaoTopOpenidConvertAPIResponse)
}

// ReleaseTaobaoTopOpenidConvertAPIResponse 将 TaobaoTopOpenidConvertAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopOpenidConvertAPIResponse(v *TaobaoTopOpenidConvertAPIResponse) {
	v.Reset()
	poolTaobaoTopOpenidConvertAPIResponse.Put(v)
}
