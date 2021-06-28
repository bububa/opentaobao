package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询物流公司信息 APIResponse
taobao.logistics.companies.get

查询淘宝网合作的物流公司信息，用于发货接口。
*/
type TaobaoLogisticsCompaniesGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLogisticsCompaniesGetResponse `json:"logistics_companies_get_response,omitempty"` 
    TaobaoLogisticsCompaniesGetResponse
}

/* model for simplify = false
type TaobaoLogisticsCompaniesGetResponse struct {

    // 物流公司信息。返回的LogisticCompany包含的具体信息为入参fields请求的字段信息。
    
    LogisticsCompanies  struct {
        LogisticsCompany  []LogisticsCompany `json:"logistics_company,omitempty"`
    } `json:"logistics_companies,omitempty"`
    

}
*/

type TaobaoLogisticsCompaniesGetResponse struct {

    // 物流公司信息。返回的LogisticCompany包含的具体信息为入参fields请求的字段信息。
    LogisticsCompanies   []LogisticsCompany `json:"logistics_companies,omitempty"`

}
