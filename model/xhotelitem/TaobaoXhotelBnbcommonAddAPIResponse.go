package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbcommonAddAPIResponse 通用调用top接口 API返回值
// taobao.xhotel.bnbcommon.add
//
// 通用调用top接口
type TaobaoXhotelBnbcommonAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbcommonAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbcommonAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbcommonAddAPIResponseModel).Reset()
}

// TaobaoXhotelBnbcommonAddAPIResponseModel is 通用调用top接口 成功返回结果
type TaobaoXhotelBnbcommonAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbcommon_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// success
	Module string `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbcommonAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = ""
}

var poolTaobaoXhotelBnbcommonAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbcommonAddAPIResponse)
	},
}

// GetTaobaoXhotelBnbcommonAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbcommonAddAPIResponse
func GetTaobaoXhotelBnbcommonAddAPIResponse() *TaobaoXhotelBnbcommonAddAPIResponse {
	return poolTaobaoXhotelBnbcommonAddAPIResponse.Get().(*TaobaoXhotelBnbcommonAddAPIResponse)
}

// ReleaseTaobaoXhotelBnbcommonAddAPIResponse 将 TaobaoXhotelBnbcommonAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbcommonAddAPIResponse(v *TaobaoXhotelBnbcommonAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbcommonAddAPIResponse.Put(v)
}
