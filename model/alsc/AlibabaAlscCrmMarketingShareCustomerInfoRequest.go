package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询分享营销客户领券信息 API请求
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

// 初始化AlibabaAlscCrmMarketingShareCustomerInfoRequest对象
func NewAlibabaAlscCrmMarketingShareCustomerInfoRequest() *AlibabaAlscCrmMarketingShareCustomerInfoRequest{
    return &AlibabaAlscCrmMarketingShareCustomerInfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.marketing.share.customer.info"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ActivityId Setter
// 活动id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetActivityId(activityId string) error {
    r.activityId = activityId
    r.Set("activity_id", activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetActivityId() string {
    return r.activityId
}
// CustomerId Setter
// 会员id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetCustomerId(customerId string) error {
    r.customerId = customerId
    r.Set("customer_id", customerId)
    return nil
}

// CustomerId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetCustomerId() string {
    return r.customerId
}
// BrandId Setter
// 品牌id(brandId和outerBrandId必传其一)
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetBrandId(brandId string) error {
    r.brandId = brandId
    r.Set("brand_id", brandId)
    return nil
}

// BrandId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetBrandId() string {
    return r.brandId
}
// OperatorId Setter
// 操作人
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOperatorId(operatorId string) error {
    r.operatorId = operatorId
    r.Set("operator_id", operatorId)
    return nil
}

// OperatorId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOperatorId() string {
    return r.operatorId
}
// OperatorName Setter
// 操作人姓名
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOperatorName(operatorName string) error {
    r.operatorName = operatorName
    r.Set("operator_name", operatorName)
    return nil
}

// OperatorName Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOperatorName() string {
    return r.operatorName
}
// OutBrandId Setter
// 外部品牌id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOutBrandId(outBrandId string) error {
    r.outBrandId = outBrandId
    r.Set("out_brand_id", outBrandId)
    return nil
}

// OutBrandId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOutBrandId() string {
    return r.outBrandId
}
// OutShopId Setter
// 外部门店id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOutShopId(outShopId string) error {
    r.outShopId = outShopId
    r.Set("out_shop_id", outShopId)
    return nil
}

// OutShopId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOutShopId() string {
    return r.outShopId
}
// RequestId Setter
// 请求幂等id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

// RequestId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetRequestId() string {
    return r.requestId
}
// ShopId Setter
// 门店id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetShopId(shopId string) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetShopId() string {
    return r.shopId
}
