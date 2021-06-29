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
    _activityId   string
    // 会员id
    _customerId   string
    // 品牌id(brandId和outerBrandId必传其一)
    _brandId   string
    // 操作人
    _operatorId   string
    // 操作人姓名
    _operatorName   string
    // 外部品牌id
    _outBrandId   string
    // 外部门店id
    _outShopId   string
    // 请求幂等id
    _requestId   string
    // 门店id
    _shopId   string
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
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetActivityId(_activityId string) error {
    r._activityId = _activityId
    r.Set("activity_id", _activityId)
    return nil
}

// ActivityId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetActivityId() string {
    return r._activityId
}
// CustomerId Setter
// 会员id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetCustomerId(_customerId string) error {
    r._customerId = _customerId
    r.Set("customer_id", _customerId)
    return nil
}

// CustomerId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetCustomerId() string {
    return r._customerId
}
// BrandId Setter
// 品牌id(brandId和outerBrandId必传其一)
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetBrandId(_brandId string) error {
    r._brandId = _brandId
    r.Set("brand_id", _brandId)
    return nil
}

// BrandId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetBrandId() string {
    return r._brandId
}
// OperatorId Setter
// 操作人
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOperatorId(_operatorId string) error {
    r._operatorId = _operatorId
    r.Set("operator_id", _operatorId)
    return nil
}

// OperatorId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOperatorId() string {
    return r._operatorId
}
// OperatorName Setter
// 操作人姓名
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOperatorName(_operatorName string) error {
    r._operatorName = _operatorName
    r.Set("operator_name", _operatorName)
    return nil
}

// OperatorName Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOperatorName() string {
    return r._operatorName
}
// OutBrandId Setter
// 外部品牌id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOutBrandId(_outBrandId string) error {
    r._outBrandId = _outBrandId
    r.Set("out_brand_id", _outBrandId)
    return nil
}

// OutBrandId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOutBrandId() string {
    return r._outBrandId
}
// OutShopId Setter
// 外部门店id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetOutShopId(_outShopId string) error {
    r._outShopId = _outShopId
    r.Set("out_shop_id", _outShopId)
    return nil
}

// OutShopId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetOutShopId() string {
    return r._outShopId
}
// RequestId Setter
// 请求幂等id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetRequestId(_requestId string) error {
    r._requestId = _requestId
    r.Set("request_id", _requestId)
    return nil
}

// RequestId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetRequestId() string {
    return r._requestId
}
// ShopId Setter
// 门店id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoRequest) SetShopId(_shopId string) error {
    r._shopId = _shopId
    r.Set("shop_id", _shopId)
    return nil
}

// ShopId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoRequest) GetShopId() string {
    return r._shopId
}
