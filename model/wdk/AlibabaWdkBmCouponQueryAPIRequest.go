package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkbmcouponqueryAPIRequest 淘鲜达券信息查询接口 API请求
// alibaba.wdk.bm.coupon.query
//
// 淘鲜达品牌营销的券信息查询接口，基于券id查询券相关信息：券id、券名称、分摊信息、面额、创建时间、开始时间、结束时间
type AlibabawdkbmcouponqueryAPIRequest struct {
	model.Params
	// 查询券参数
	_isvQueryCouponParam *IsvQueryCouponParam
}

// NewAlibabawdkbmcouponqueryRequest 初始化AlibabawdkbmcouponqueryAPIRequest对象
func NewAlibabawdkbmcouponqueryRequest() *AlibabawdkbmcouponqueryAPIRequest {
	return &AlibabawdkbmcouponqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkbmcouponqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.coupon.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkbmcouponqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkbmcouponqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvQueryCouponParam is IsvQueryCouponParam Setter
// 查询券参数
func (r *AlibabawdkbmcouponqueryAPIRequest) SetIsvQueryCouponParam(_isvQueryCouponParam *IsvQueryCouponParam) error {
	r._isvQueryCouponParam = _isvQueryCouponParam
	r.Set("isv_query_coupon_param", _isvQueryCouponParam)
	return nil
}

// GetIsvQueryCouponParam IsvQueryCouponParam Getter
func (r AlibabawdkbmcouponqueryAPIRequest) GetIsvQueryCouponParam() *IsvQueryCouponParam {
	return r._isvQueryCouponParam
}
