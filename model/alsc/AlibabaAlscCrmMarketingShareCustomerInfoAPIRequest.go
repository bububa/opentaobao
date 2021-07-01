package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest
查询分享营销客户领券信息 API请求
alibaba.alsc.crm.marketing.share.customer.info

查询分享营销活动的客户领券信息 */
type AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest struct {
	model.Params
	// 活动id
	_activityId string
	// 会员id
	_customerId string
	// 品牌id(brandId和outerBrandId必传其一)
	_brandId string
	// 操作人
	_operatorId string
	// 操作人姓名
	_operatorName string
	// 外部品牌id
	_outBrandId string
	// 外部门店id
	_outShopId string
	// 请求幂等id
	_requestId string
	// 门店id
	_shopId string
}

// NewAlibabaAlscCrmMarketingShareCustomerInfoRequest 初始化AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest对象
func NewAlibabaAlscCrmMarketingShareCustomerInfoRequest() *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest {
	return &AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.marketing.share.customer.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActivityId Setter
// 活动id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetActivityId() string {
	return r._activityId
}

// Set is CustomerId Setter
// 会员id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) SetCustomerId(_customerId string) error {
	r._customerId = _customerId
	r.Set("customer_id", _customerId)
	return nil
}

// Get CustomerId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetCustomerId() string {
	return r._customerId
}

// Set is BrandId Setter
// 品牌id(brandId和outerBrandId必传其一)
func (r *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) SetBrandId(_brandId string) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// Get BrandId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetBrandId() string {
	return r._brandId
}

// Set is OperatorId Setter
// 操作人
func (r *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) SetOperatorId(_operatorId string) error {
	r._operatorId = _operatorId
	r.Set("operator_id", _operatorId)
	return nil
}

// Get OperatorId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetOperatorId() string {
	return r._operatorId
}

// Set is OperatorName Setter
// 操作人姓名
func (r *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) SetOperatorName(_operatorName string) error {
	r._operatorName = _operatorName
	r.Set("operator_name", _operatorName)
	return nil
}

// Get OperatorName Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetOperatorName() string {
	return r._operatorName
}

// Set is OutBrandId Setter
// 外部品牌id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) SetOutBrandId(_outBrandId string) error {
	r._outBrandId = _outBrandId
	r.Set("out_brand_id", _outBrandId)
	return nil
}

// Get OutBrandId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetOutBrandId() string {
	return r._outBrandId
}

// Set is OutShopId Setter
// 外部门店id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) SetOutShopId(_outShopId string) error {
	r._outShopId = _outShopId
	r.Set("out_shop_id", _outShopId)
	return nil
}

// Get OutShopId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetOutShopId() string {
	return r._outShopId
}

// Set is RequestId Setter
// 请求幂等id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// Get RequestId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetRequestId() string {
	return r._requestId
}

// Set is ShopId Setter
// 门店id
func (r *AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// Get ShopId Getter
func (r AlibabaAlscCrmMarketingShareCustomerInfoAPIRequest) GetShopId() string {
	return r._shopId
}
