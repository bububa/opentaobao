package cainiaocntec

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest 验证码验证 API请求
// cainiao.cntec.locallife.communitylife.verifyservicecode
//
// 验证码验证
type CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest struct {
	model.Params
	// 订单id
	_orderNo string
	// 扩展信息Json格式
	_feature string
	// 服务码
	_serviceCode string
}

// NewCainiaocnteclocallifecommunitylifeverifyservicecodeRequest 初始化CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest对象
func NewCainiaocnteclocallifecommunitylifeverifyservicecodeRequest() *CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest {
	return &CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.locallife.communitylife.verifyservicecode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 订单id
func (r *CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetFeature is Feature Setter
// 扩展信息Json格式
func (r *CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest) GetFeature() string {
	return r._feature
}

// SetServiceCode is ServiceCode Setter
// 服务码
func (r *CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r CainiaocnteclocallifecommunitylifeverifyservicecodeAPIRequest) GetServiceCode() string {
	return r._serviceCode
}
