package cainiaocntec

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest 验证码验证 API请求
// cainiao.cntec.locallife.communitylife.verifyservicecode
//
// 验证码验证
type CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest struct {
	model.Params
	// 订单id
	_orderNo string
	// 扩展信息Json格式
	_feature string
	// 服务码
	_serviceCode string
}

// NewCainiaoCntecLocallifeCommunitylifeVerifyservicecodeRequest 初始化CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest对象
func NewCainiaoCntecLocallifeCommunitylifeVerifyservicecodeRequest() *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest {
	return &CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.locallife.communitylife.verifyservicecode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderNo is OrderNo Setter
// 订单id
func (r *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetFeature is Feature Setter
// 扩展信息Json格式
func (r *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) GetFeature() string {
	return r._feature
}

// SetServiceCode is ServiceCode Setter
// 服务码
func (r *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) GetServiceCode() string {
	return r._serviceCode
}
