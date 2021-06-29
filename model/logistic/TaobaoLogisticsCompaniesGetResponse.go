package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询物流公司信息 API返回值 
taobao.logistics.companies.get

查询淘宝网合作的物流公司信息，用于发货接口。
*/
type TaobaoLogisticsCompaniesGetAPIResponse struct {
    model.CommonResponse
    TaobaoLogisticsCompaniesGetResponse
}

// 查询物流公司信息 成功返回结果
type TaobaoLogisticsCompaniesGetResponse struct {
    XMLName xml.Name `xml:"logistics_companies_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 物流公司信息。返回的LogisticCompany包含的具体信息为入参fields请求的字段信息。
    LogisticsCompanies   []LogisticsCompany `json:"logistics_companies,omitempty" xml:"logistics_companies>logistics_company,omitempty"`
}
