package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBaseinfoGetAPIResponse 酒店基础信息查询接口 API返回值
// taobao.xhotel.baseinfo.get
//
// 酒店基础信息(酒店/房型/房价定义)查询接口， 包括 酒店房型可售, 以及 hid 下 的标准房型列表
type TaobaoXhotelBaseinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBaseinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBaseinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBaseinfoGetAPIResponseModel).Reset()
}

// TaobaoXhotelBaseinfoGetAPIResponseModel is 酒店基础信息查询接口 成功返回结果
type TaobaoXhotelBaseinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_baseinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelBaseinfoGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBaseinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelBaseinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBaseinfoGetAPIResponse)
	},
}

// GetTaobaoXhotelBaseinfoGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelBaseinfoGetAPIResponse
func GetTaobaoXhotelBaseinfoGetAPIResponse() *TaobaoXhotelBaseinfoGetAPIResponse {
	return poolTaobaoXhotelBaseinfoGetAPIResponse.Get().(*TaobaoXhotelBaseinfoGetAPIResponse)
}

// ReleaseTaobaoXhotelBaseinfoGetAPIResponse 将 TaobaoXhotelBaseinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBaseinfoGetAPIResponse(v *TaobaoXhotelBaseinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBaseinfoGetAPIResponse.Put(v)
}
