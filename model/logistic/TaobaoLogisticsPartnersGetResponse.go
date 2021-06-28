package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询支持起始地到目的地范围的物流公司 APIResponse
taobao.logistics.partners.get

查询物流公司信息（可以查询目的地可不可达情况）
*/
type TaobaoLogisticsPartnersGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoLogisticsPartnersGetResponse `json:"logistics_partners_get_response,omitempty"` 
    TaobaoLogisticsPartnersGetResponse
}

/* model for simplify = false
type TaobaoLogisticsPartnersGetResponse struct {

    // 查询揽送范围之内的物流公司信息
    
    LogisticsPartners  struct {
        LogisticsPartner  []LogisticsPartner `json:"logistics_partner,omitempty"`
    } `json:"logistics_partners,omitempty"`
    

}
*/

type TaobaoLogisticsPartnersGetResponse struct {

    // 查询揽送范围之内的物流公司信息
    LogisticsPartners   []LogisticsPartner `json:"logistics_partners,omitempty"`

}
