package waybill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoWaybillIiSearchAPIResponse 查询面单服务订购及面单使用情况 API返回值
// cainiao.waybill.ii.search
//
// 获取发货地&amp;CP开通状态&amp;账户的使用情况
type CainiaoWaybillIiSearchAPIResponse struct {
	model.CommonResponse
	CainiaoWaybillIiSearchAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoWaybillIiSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoWaybillIiSearchAPIResponseModel).Reset()
}

// CainiaoWaybillIiSearchAPIResponseModel is 查询面单服务订购及面单使用情况 成功返回结果
type CainiaoWaybillIiSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_waybill_ii_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// CP网点信息及对应的商家的发货信息
	WaybillApplySubscriptionCols []WaybillApplySubscriptionInfo `json:"waybill_apply_subscription_cols,omitempty" xml:"waybill_apply_subscription_cols>waybill_apply_subscription_info,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoWaybillIiSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WaybillApplySubscriptionCols = m.WaybillApplySubscriptionCols[:0]
}

var poolCainiaoWaybillIiSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoWaybillIiSearchAPIResponse)
	},
}

// GetCainiaoWaybillIiSearchAPIResponse 从 sync.Pool 获取 CainiaoWaybillIiSearchAPIResponse
func GetCainiaoWaybillIiSearchAPIResponse() *CainiaoWaybillIiSearchAPIResponse {
	return poolCainiaoWaybillIiSearchAPIResponse.Get().(*CainiaoWaybillIiSearchAPIResponse)
}

// ReleaseCainiaoWaybillIiSearchAPIResponse 将 CainiaoWaybillIiSearchAPIResponse 保存到 sync.Pool
func ReleaseCainiaoWaybillIiSearchAPIResponse(v *CainiaoWaybillIiSearchAPIResponse) {
	v.Reset()
	poolCainiaoWaybillIiSearchAPIResponse.Put(v)
}
