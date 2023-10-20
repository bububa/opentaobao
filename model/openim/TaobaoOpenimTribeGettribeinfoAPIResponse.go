package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeGettribeinfoAPIResponse 获取群信息 API返回值
// taobao.openim.tribe.gettribeinfo
//
// 获取群信息
type TaobaoOpenimTribeGettribeinfoAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeGettribeinfoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeGettribeinfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeGettribeinfoAPIResponseModel).Reset()
}

// TaobaoOpenimTribeGettribeinfoAPIResponseModel is 获取群信息 成功返回结果
type TaobaoOpenimTribeGettribeinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_gettribeinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群信息
	TribeInfo *TribeInfo `json:"tribe_info,omitempty" xml:"tribe_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeGettribeinfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeInfo = nil
}

var poolTaobaoOpenimTribeGettribeinfoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeGettribeinfoAPIResponse)
	},
}

// GetTaobaoOpenimTribeGettribeinfoAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeGettribeinfoAPIResponse
func GetTaobaoOpenimTribeGettribeinfoAPIResponse() *TaobaoOpenimTribeGettribeinfoAPIResponse {
	return poolTaobaoOpenimTribeGettribeinfoAPIResponse.Get().(*TaobaoOpenimTribeGettribeinfoAPIResponse)
}

// ReleaseTaobaoOpenimTribeGettribeinfoAPIResponse 将 TaobaoOpenimTribeGettribeinfoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeGettribeinfoAPIResponse(v *TaobaoOpenimTribeGettribeinfoAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeGettribeinfoAPIResponse.Put(v)
}
