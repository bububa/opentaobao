package alitripmerchant

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest 抽奖用户名单查询接口 API请求
// alitrip.merchant.galaxy.wechat.data.lottery.query
//
// 抽奖用户名单查询接口
type AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 查询抽奖用户数据入参
	_queryLotteryDataDTO *QueryLotteryDataDto
}

// NewAlitripmerchantgalaxywechatdatalotteryqueryRequest 初始化AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest对象
func NewAlitripmerchantgalaxywechatdatalotteryqueryRequest() *AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest {
	return &AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.data.lottery.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetQueryLotteryDataDTO is QueryLotteryDataDTO Setter
// 查询抽奖用户数据入参
func (r *AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest) SetQueryLotteryDataDTO(_queryLotteryDataDTO *QueryLotteryDataDto) error {
	r._queryLotteryDataDTO = _queryLotteryDataDTO
	r.Set("query_lottery_data_d_t_o", _queryLotteryDataDTO)
	return nil
}

// GetQueryLotteryDataDTO QueryLotteryDataDTO Getter
func (r AlitripmerchantgalaxywechatdatalotteryqueryAPIRequest) GetQueryLotteryDataDTO() *QueryLotteryDataDto {
	return r._queryLotteryDataDTO
}
