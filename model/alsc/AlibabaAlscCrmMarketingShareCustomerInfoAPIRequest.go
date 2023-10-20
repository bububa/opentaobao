package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmmarketingsharecustomerinfoAPIRequest 查询分享营销客户领券信息 API请求
// alibaba.alsc.crm.marketing.share.customer.info
//
// 查询分享营销活动的客户领券信息
type AlibabaalsccrmmarketingsharecustomerinfoAPIRequest struct {
	model.Params
	// 活动id
	_activityId string
	// 品牌id(brandId和outerBrandId必传其一)
	_brandId string
	// 会员id
	_customerId string
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

// NewAlibabaalsccrmmarketingsharecustomerinfoRequest 初始化AlibabaalsccrmmarketingsharecustomerinfoAPIRequest对象
func NewAlibabaalsccrmmarketingsharecustomerinfoRequest() *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest {
	return &AlibabaalsccrmmarketingsharecustomerinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.marketing.share.customer.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) SetActivityId(_activityId string) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetActivityId() string {
	return r._activityId
}

// SetBrandId is BrandId Setter
// 品牌id(brandId和outerBrandId必传其一)
func (r *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) SetBrandId(_brandId string) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetBrandId() string {
	return r._brandId
}

// SetCustomerId is CustomerId Setter
// 会员id
func (r *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) SetCustomerId(_customerId string) error {
	r._customerId = _customerId
	r.Set("customer_id", _customerId)
	return nil
}

// GetCustomerId CustomerId Getter
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetCustomerId() string {
	return r._customerId
}

// SetOperatorId is OperatorId Setter
// 操作人
func (r *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) SetOperatorId(_operatorId string) error {
	r._operatorId = _operatorId
	r.Set("operator_id", _operatorId)
	return nil
}

// GetOperatorId OperatorId Getter
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetOperatorId() string {
	return r._operatorId
}

// SetOperatorName is OperatorName Setter
// 操作人姓名
func (r *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) SetOperatorName(_operatorName string) error {
	r._operatorName = _operatorName
	r.Set("operator_name", _operatorName)
	return nil
}

// GetOperatorName OperatorName Getter
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetOperatorName() string {
	return r._operatorName
}

// SetOutBrandId is OutBrandId Setter
// 外部品牌id
func (r *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) SetOutBrandId(_outBrandId string) error {
	r._outBrandId = _outBrandId
	r.Set("out_brand_id", _outBrandId)
	return nil
}

// GetOutBrandId OutBrandId Getter
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetOutBrandId() string {
	return r._outBrandId
}

// SetOutShopId is OutShopId Setter
// 外部门店id
func (r *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) SetOutShopId(_outShopId string) error {
	r._outShopId = _outShopId
	r.Set("out_shop_id", _outShopId)
	return nil
}

// GetOutShopId OutShopId Getter
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetOutShopId() string {
	return r._outShopId
}

// SetRequestId is RequestId Setter
// 请求幂等id
func (r *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetShopId is ShopId Setter
// 门店id
func (r *AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabaalsccrmmarketingsharecustomerinfoAPIRequest) GetShopId() string {
	return r._shopId
}
