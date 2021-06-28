package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子面单物流详情授权url获取 APIResponse
cainiao.waybill.ii.logisticsdetail.url.get

获取电子面单物流详情授权访问的H5 url
*/
type CainiaoWaybillIiLogisticsdetailUrlGetAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoWaybillIiLogisticsdetailUrlGetResponse `json:"cainiao_waybill_ii_logisticsdetail_url_get_response,omitempty"` 
    CainiaoWaybillIiLogisticsdetailUrlGetResponse
}

/* model for simplify = false
type CainiaoWaybillIiLogisticsdetailUrlGetResponse struct {

    // 授权访问的url
    
    Url   string `json:"url,omitempty"`
    

}
*/

type CainiaoWaybillIiLogisticsdetailUrlGetResponse struct {

    // 授权访问的url
    Url   string `json:"url,omitempty"`

}
