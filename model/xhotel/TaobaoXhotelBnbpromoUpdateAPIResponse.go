package xhotel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbpromoUpdateAPIResponse 民宿营销活动更新 API返回值
// taobao.xhotel.bnbpromo.update
//
// 全量更新对应外部活动code相关的营销活动信息
type TaobaoXhotelBnbpromoUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelBnbpromoUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelBnbpromoUpdateAPIResponseModel).Reset()
}

// TaobaoXhotelBnbpromoUpdateAPIResponseModel is 民宿营销活动更新 成功返回结果
type TaobaoXhotelBnbpromoUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_bnbpromo_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果集
	Result *TaobaoXhotelBnbpromoUpdateResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelBnbpromoUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelBnbpromoUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelBnbpromoUpdateAPIResponse)
	},
}

// GetTaobaoXhotelBnbpromoUpdateAPIResponse 从 sync.Pool 获取 TaobaoXhotelBnbpromoUpdateAPIResponse
func GetTaobaoXhotelBnbpromoUpdateAPIResponse() *TaobaoXhotelBnbpromoUpdateAPIResponse {
	return poolTaobaoXhotelBnbpromoUpdateAPIResponse.Get().(*TaobaoXhotelBnbpromoUpdateAPIResponse)
}

// ReleaseTaobaoXhotelBnbpromoUpdateAPIResponse 将 TaobaoXhotelBnbpromoUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelBnbpromoUpdateAPIResponse(v *TaobaoXhotelBnbpromoUpdateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelBnbpromoUpdateAPIResponse.Put(v)
}
