package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoGetAPIResponse 民宿查询营销活动 API返回值
// taobao.xhotel.bnbpromo.get
//
// 民宿查询营销活动
type TaobaoXhotelBnbpromoGetAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbpromoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbpromoGetAPIResponseModel).Reset()
}

// TaobaoXhotelBnbpromoGetAPIResponseModel is 民宿查询营销活动 成功返回结果
type TaobaoXhotelBnbpromoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelBnbpromoGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelBnbpromoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbpromoGetAPIResponse)
	},
}

// GetTaobaoXhotelBnbpromoGetAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbpromoGetAPIResponse
func GetTaobaoXhotelBnbpromoGetAPIResponse() *TaobaoXhotelBnbpromoGetAPIResponse {
	return poolTaobaoXhotelBnbpromoGetAPIResponse.Get().(*TaobaoXhotelBnbpromoGetAPIResponse)
}

// ReleaseTaobaoXhotelBnbpromoGetAPIResponse 将 TaobaoXhotelBnbpromoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbpromoGetAPIResponse(v *TaobaoXhotelBnbpromoGetAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbpromoGetAPIResponse.Put(v)
}
