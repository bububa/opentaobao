package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMultiplerateDeleteAPIResponse 复杂价格删除 API返回值
// taobao.xhotel.multiplerate.delete
//
// 酒店产品库rate删除
type TaobaoXhotelMultiplerateDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelMultiplerateDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelMultiplerateDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelMultiplerateDeleteAPIResponseModel).Reset()
}

// TaobaoXhotelMultiplerateDeleteAPIResponseModel is 复杂价格删除 成功返回结果
type TaobaoXhotelMultiplerateDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_multiplerate_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelMultiplerateDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelMultiplerateDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelMultiplerateDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelMultiplerateDeleteAPIResponse)
	},
}

// GetTaobaoXhotelMultiplerateDeleteAPIResponse 从 sync.Pool 获取 TaobaoXhotelMultiplerateDeleteAPIResponse
func GetTaobaoXhotelMultiplerateDeleteAPIResponse() *TaobaoXhotelMultiplerateDeleteAPIResponse {
	return poolTaobaoXhotelMultiplerateDeleteAPIResponse.Get().(*TaobaoXhotelMultiplerateDeleteAPIResponse)
}

// ReleaseTaobaoXhotelMultiplerateDeleteAPIResponse 将 TaobaoXhotelMultiplerateDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelMultiplerateDeleteAPIResponse(v *TaobaoXhotelMultiplerateDeleteAPIResponse) {
	v.Reset()
	poolTaobaoXhotelMultiplerateDeleteAPIResponse.Put(v)
}
