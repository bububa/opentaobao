package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询能否代扣 APIRequest
cainiao.endpoint.locker.top.withhold.query

查询是否有代扣欠款，是否签署代扣协议。
*/
type CainiaoEndpointLockerTopWithholdQueryRequest struct {
    model.Params

    // 柜子公司编码
    companyCode   string 

    // 开放用户Id
    openUserId   string 

    // 柜子业务：0-取件业务，1-寄件业务，2-派样业务，3-小件员约柜（在约期内仅能使用一次）4-小件员租柜（在约期内可反复使用）
    orderType   int64 

}

func NewCainiaoEndpointLockerTopWithholdQueryRequest() *CainiaoEndpointLockerTopWithholdQueryRequest{
    return &CainiaoEndpointLockerTopWithholdQueryRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoEndpointLockerTopWithholdQueryRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.withhold.query"
}

func (r CainiaoEndpointLockerTopWithholdQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoEndpointLockerTopWithholdQueryRequest) SetCompanyCode(companyCode string) error {
    r.companyCode = companyCode
    r.Set("company_code", companyCode)
    return nil
}

func (r CainiaoEndpointLockerTopWithholdQueryRequest) GetCompanyCode() string {
    return r.companyCode
}

func (r *CainiaoEndpointLockerTopWithholdQueryRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

func (r CainiaoEndpointLockerTopWithholdQueryRequest) GetOpenUserId() string {
    return r.openUserId
}

func (r *CainiaoEndpointLockerTopWithholdQueryRequest) SetOrderType(orderType int64) error {
    r.orderType = orderType
    r.Set("order_type", orderType)
    return nil
}

func (r CainiaoEndpointLockerTopWithholdQueryRequest) GetOrderType() int64 {
    return r.orderType
}

