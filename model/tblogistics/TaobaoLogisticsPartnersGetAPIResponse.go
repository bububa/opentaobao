package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsPartnersGetAPIResponse 查询支持起始地到目的地范围的物流公司 API返回值
// taobao.logistics.partners.get
//
// 查询物流公司信息（可以查询目的地可不可达情况）
type TaobaoLogisticsPartnersGetAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsPartnersGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsPartnersGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsPartnersGetAPIResponseModel).Reset()
}

// TaobaoLogisticsPartnersGetAPIResponseModel is 查询支持起始地到目的地范围的物流公司 成功返回结果
type TaobaoLogisticsPartnersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_partners_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询揽送范围之内的物流公司信息
	LogisticsPartners []LogisticsPartner `json:"logistics_partners,omitempty" xml:"logistics_partners>logistics_partner,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsPartnersGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.LogisticsPartners = m.LogisticsPartners[:0]
}

var poolTaobaoLogisticsPartnersGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsPartnersGetAPIResponse)
	},
}

// GetTaobaoLogisticsPartnersGetAPIResponse 从 sync.Pool 获取 TaobaoLogisticsPartnersGetAPIResponse
func GetTaobaoLogisticsPartnersGetAPIResponse() *TaobaoLogisticsPartnersGetAPIResponse {
	return poolTaobaoLogisticsPartnersGetAPIResponse.Get().(*TaobaoLogisticsPartnersGetAPIResponse)
}

// ReleaseTaobaoLogisticsPartnersGetAPIResponse 将 TaobaoLogisticsPartnersGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsPartnersGetAPIResponse(v *TaobaoLogisticsPartnersGetAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsPartnersGetAPIResponse.Put(v)
}
