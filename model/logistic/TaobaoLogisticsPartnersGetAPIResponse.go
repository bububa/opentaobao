package logistic

import (
	"encoding/xml"

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

// TaobaoLogisticsPartnersGetAPIResponseModel is 查询支持起始地到目的地范围的物流公司 成功返回结果
type TaobaoLogisticsPartnersGetAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_partners_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询揽送范围之内的物流公司信息
	LogisticsPartners []LogisticsPartner `json:"logistics_partners,omitempty" xml:"logistics_partners>logistics_partner,omitempty"`
}
