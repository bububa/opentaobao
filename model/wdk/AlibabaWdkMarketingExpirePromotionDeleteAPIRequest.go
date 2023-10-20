package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingexpirepromotiondeleteAPIRequest 短保优惠删除 API请求
// alibaba.wdk.marketing.expire.promotion.delete
//
// 短保优惠删除
type AlibabawdkmarketingexpirepromotiondeleteAPIRequest struct {
	model.Params
	// 删除短保优惠
	_param0 *ExpirePromotionBo
}

// NewAlibabawdkmarketingexpirepromotiondeleteRequest 初始化AlibabawdkmarketingexpirepromotiondeleteAPIRequest对象
func NewAlibabawdkmarketingexpirepromotiondeleteRequest() *AlibabawdkmarketingexpirepromotiondeleteAPIRequest {
	return &AlibabawdkmarketingexpirepromotiondeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingexpirepromotiondeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.expire.promotion.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingexpirepromotiondeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingexpirepromotiondeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 删除短保优惠
func (r *AlibabawdkmarketingexpirepromotiondeleteAPIRequest) SetParam0(_param0 *ExpirePromotionBo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabawdkmarketingexpirepromotiondeleteAPIRequest) GetParam0() *ExpirePromotionBo {
	return r._param0
}
