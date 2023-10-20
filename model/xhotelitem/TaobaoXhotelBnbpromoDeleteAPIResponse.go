package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoDeleteAPIResponse 民宿卖家活动删除 API返回值
// taobao.xhotel.bnbpromo.delete
//
// 民宿删除营销活动
type TaobaoXhotelBnbpromoDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbpromoDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbpromoDeleteAPIResponseModel).Reset()
}

// TaobaoXhotelBnbpromoDeleteAPIResponseModel is 民宿卖家活动删除 成功返回结果
type TaobaoXhotelBnbpromoDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelBnbpromoDeleteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelBnbpromoDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbpromoDeleteAPIResponse)
	},
}

// GetTaobaoXhotelBnbpromoDeleteAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbpromoDeleteAPIResponse
func GetTaobaoXhotelBnbpromoDeleteAPIResponse() *TaobaoXhotelBnbpromoDeleteAPIResponse {
	return poolTaobaoXhotelBnbpromoDeleteAPIResponse.Get().(*TaobaoXhotelBnbpromoDeleteAPIResponse)
}

// ReleaseTaobaoXhotelBnbpromoDeleteAPIResponse 将 TaobaoXhotelBnbpromoDeleteAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbpromoDeleteAPIResponse(v *TaobaoXhotelBnbpromoDeleteAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbpromoDeleteAPIResponse.Put(v)
}
