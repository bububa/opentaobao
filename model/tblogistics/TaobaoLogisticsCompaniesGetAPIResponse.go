package tblogistics

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsCompaniesGetAPIResponse 查询物流公司信息 API返回值
// taobao.logistics.companies.get
//
// 查询淘宝网合作的物流公司信息，用于发货接口。
type TaobaoLogisticsCompaniesGetAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsCompaniesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsCompaniesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsCompaniesGetAPIResponseModel).Reset()
}

// TaobaoLogisticsCompaniesGetAPIResponseModel is 查询物流公司信息 成功返回结果
type TaobaoLogisticsCompaniesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_companies_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 物流公司信息。返回的LogisticCompany包含的具体信息为入参fields请求的字段信息。
	LogisticsCompanies []LogisticsCompany `json:"logistics_companies,omitempty" xml:"logistics_companies>logistics_company,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsCompaniesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.LogisticsCompanies = m.LogisticsCompanies[:0]
}

var poolTaobaoLogisticsCompaniesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsCompaniesGetAPIResponse)
	},
}

// GetTaobaoLogisticsCompaniesGetAPIResponse 从 sync.Pool 获取 TaobaoLogisticsCompaniesGetAPIResponse
func GetTaobaoLogisticsCompaniesGetAPIResponse() *TaobaoLogisticsCompaniesGetAPIResponse {
	return poolTaobaoLogisticsCompaniesGetAPIResponse.Get().(*TaobaoLogisticsCompaniesGetAPIResponse)
}

// ReleaseTaobaoLogisticsCompaniesGetAPIResponse 将 TaobaoLogisticsCompaniesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsCompaniesGetAPIResponse(v *TaobaoLogisticsCompaniesGetAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsCompaniesGetAPIResponse.Put(v)
}
