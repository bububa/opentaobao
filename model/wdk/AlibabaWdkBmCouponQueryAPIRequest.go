package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkBmCouponQueryAPIRequest) Reset() {
	r._isvQueryCouponParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmCouponQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.coupon.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmCouponQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkBmCouponQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkBmCouponQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkBmCouponQueryRequest()
	},
}

// GetAlibabaWdkBmCouponQueryRequest 从 sync.Pool 获取 AlibabaWdkBmCouponQueryAPIRequest
func GetAlibabaWdkBmCouponQueryAPIRequest() *AlibabaWdkBmCouponQueryAPIRequest {
	return poolAlibabaWdkBmCouponQueryAPIRequest.Get().(*AlibabaWdkBmCouponQueryAPIRequest)
}

// ReleaseAlibabaWdkBmCouponQueryAPIRequest 将 AlibabaWdkBmCouponQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkBmCouponQueryAPIRequest(v *AlibabaWdkBmCouponQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkBmCouponQueryAPIRequest.Put(v)
}
