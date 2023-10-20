package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbownerAddAPIResponse 民宿房东信息添加 API返回值
// taobao.xhotel.bnbowner.add
//
// 添加和更新民宿房东的信息
type TaobaoXhotelBnbownerAddAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbownerAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbownerAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbownerAddAPIResponseModel).Reset()
}

// TaobaoXhotelBnbownerAddAPIResponseModel is 民宿房东信息添加 成功返回结果
type TaobaoXhotelBnbownerAddAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbowner_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelBnbownerAddResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbownerAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelBnbownerAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbownerAddAPIResponse)
	},
}

// GetTaobaoXhotelBnbownerAddAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbownerAddAPIResponse
func GetTaobaoXhotelBnbownerAddAPIResponse() *TaobaoXhotelBnbownerAddAPIResponse {
	return poolTaobaoXhotelBnbownerAddAPIResponse.Get().(*TaobaoXhotelBnbownerAddAPIResponse)
}

// ReleaseTaobaoXhotelBnbownerAddAPIResponse 将 TaobaoXhotelBnbownerAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbownerAddAPIResponse(v *TaobaoXhotelBnbownerAddAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbownerAddAPIResponse.Put(v)
}
