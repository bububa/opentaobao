package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelServicetimeGetAPIResponse 查询实体对应的服务时间数据 API返回值
// taobao.xhotel.servicetime.get
//
// 通过实体来获取对应的服务时间数据
type TaobaoXhotelServicetimeGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelServicetimeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelServicetimeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelServicetimeGetAPIResponseModel).Reset()
}

// TaobaoXhotelServicetimeGetAPIResponseModel is 查询实体对应的服务时间数据 成功返回结果
type TaobaoXhotelServicetimeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_servicetime_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelServicetimeGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelServicetimeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelServicetimeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelServicetimeGetAPIResponse)
	},
}

// GetTaobaoXhotelServicetimeGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelServicetimeGetAPIResponse
func GetTaobaoXhotelServicetimeGetAPIResponse() *TaobaoXhotelServicetimeGetAPIResponse {
	return poolTaobaoXhotelServicetimeGetAPIResponse.Get().(*TaobaoXhotelServicetimeGetAPIResponse)
}

// ReleaseTaobaoXhotelServicetimeGetAPIResponse 将 TaobaoXhotelServicetimeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelServicetimeGetAPIResponse(v *TaobaoXhotelServicetimeGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelServicetimeGetAPIResponse.Put(v)
}
