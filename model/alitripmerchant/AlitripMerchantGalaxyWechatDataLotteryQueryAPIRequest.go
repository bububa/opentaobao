package alitripmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest 抽奖用户名单查询接口 API请求
// alitrip.merchant.galaxy.wechat.data.lottery.query
//
// 抽奖用户名单查询接口
type AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest struct {
	model.Params
	// 租户标识
	_tenantKey string
	// 查询抽奖用户数据入参
	_queryLotteryDataDTO *QueryLotteryDataDto
}

// NewAlitripMerchantGalaxyWechatDataLotteryQueryRequest 初始化AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest对象
func NewAlitripMerchantGalaxyWechatDataLotteryQueryRequest() *AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest {
	return &AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest) Reset() {
	r._tenantKey = ""
	r._queryLotteryDataDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.merchant.galaxy.wechat.data.lottery.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTenantKey is TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest) SetTenantKey(_tenantKey string) error {
	r._tenantKey = _tenantKey
	r.Set("tenant_key", _tenantKey)
	return nil
}

// GetTenantKey TenantKey Getter
func (r AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest) GetTenantKey() string {
	return r._tenantKey
}

// SetQueryLotteryDataDTO is QueryLotteryDataDTO Setter
// 查询抽奖用户数据入参
func (r *AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest) SetQueryLotteryDataDTO(_queryLotteryDataDTO *QueryLotteryDataDto) error {
	r._queryLotteryDataDTO = _queryLotteryDataDTO
	r.Set("query_lottery_data_d_t_o", _queryLotteryDataDTO)
	return nil
}

// GetQueryLotteryDataDTO QueryLotteryDataDTO Getter
func (r AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest) GetQueryLotteryDataDTO() *QueryLotteryDataDto {
	return r._queryLotteryDataDTO
}

var poolAlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripMerchantGalaxyWechatDataLotteryQueryRequest()
	},
}

// GetAlitripMerchantGalaxyWechatDataLotteryQueryRequest 从 sync.Pool 获取 AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest
func GetAlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest() *AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest {
	return poolAlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest.Get().(*AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest)
}

// ReleaseAlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest 将 AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest(v *AlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest) {
	v.Reset()
	poolAlitripMerchantGalaxyWechatDataLotteryQueryAPIRequest.Put(v)
}
