package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmCouponQueryAPIRequest 淘鲜达券信息查询接口 API请求
// alibaba.wdk.bm.coupon.query
//
// 淘鲜达品牌营销的券信息查询接口，基于券id查询券相关信息：券id、券名称、分摊信息、面额、创建时间、开始时间、结束时间
type AlibabaWdkBmCouponQueryAPIRequest struct {
	model.Params
	// 查询券参数
	_isvQueryCouponParam *IsvQueryCouponParam
}

// NewAlibabaWdkBmCouponQueryRequest 初始化AlibabaWdkBmCouponQueryAPIRequest对象
func NewAlibabaWdkBmCouponQueryRequest() *AlibabaWdkBmCouponQueryAPIRequest {
	return &AlibabaWdkBmCouponQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmCouponQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.coupon.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmCouponQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIsvQueryCouponParam is IsvQueryCouponParam Setter
// 查询券参数
func (r *AlibabaWdkBmCouponQueryAPIRequest) SetIsvQueryCouponParam(_isvQueryCouponParam *IsvQueryCouponParam) error {
	r._isvQueryCouponParam = _isvQueryCouponParam
	r.Set("isv_query_coupon_param", _isvQueryCouponParam)
	return nil
}

// GetIsvQueryCouponParam IsvQueryCouponParam Getter
func (r AlibabaWdkBmCouponQueryAPIRequest) GetIsvQueryCouponParam() *IsvQueryCouponParam {
	return r._isvQueryCouponParam
}
