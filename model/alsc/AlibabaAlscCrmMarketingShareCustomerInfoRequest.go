package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分享营销客户领券信息 APIRequest
alibaba.alsc.crm.marketing.share.customer.info

查询分享营销活动的客户领券信息
*/
type AlibabaAlscCrmMarketingShareCustomerInfoRequest struct {
    model.Params

    // 活动id
    activityId   string 

    // 会员id
    customerId   string 

    // 品牌id(brandId和outerBrandId必传其一)
    brandId   string 

    // 操作人
    operatorId   string 

    // 操作人姓名
    operatorName   string 

    // 外部品牌id
    outBrandId   string 

    // 外部门店id
    outShopId   string 

    // 请求幂等id
    requestId   string 

    // 门店id
    shopId   string 

}

func NewAlibabaAlscCrmMarketingShareCustomerInfoRequest() *AlibabaAlscCrmMarketingShareCustomerInfoRequest{
    return &AlibabaAlscCrmMarketingShareCustomerInfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.marketing.share.customer.info"
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetActivityId() string {
    return r.activityId
}

func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetCustomerId(customerId string) error {
    r.customerId = customerId
    r.Set("customer_id", customerId)
    return nil
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetCustomerId() string {
    return r.customerId
}

func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetBrandId(brandId string) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetBrandId() string {
    return r.brandId
}

func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOperatorId(operatorId string) error {
    r.operatorId = operatorId
    r.Set("operator_id", operatorId)
    return nil
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOperatorId() string {
    return r.operatorId
}

func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOperatorName(operatorName string) error {
    r.operatorName = operatorName
    r.Set("operator_name", operatorName)
    return nil
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOperatorName() string {
    return r.operatorName
}

func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOutBrandId(outBrandId string) error {
    r.outBrandId = outBrandId
    r.Set("out_brand_id", outBrandId)
    return nil
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOutBrandId() string {
    return r.outBrandId
}

func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOutShopId(outShopId string) error {
    r.outShopId = outShopId
    r.Set("out_shop_id", outShopId)
    return nil
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOutShopId() string {
    return r.outShopId
}

func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetRequestId() string {
    return r.requestId
}

func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetShopId() string {
    return r.shopId
}

