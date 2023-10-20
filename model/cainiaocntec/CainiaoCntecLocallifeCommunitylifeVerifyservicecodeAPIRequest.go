package cainiaocntec

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) Reset() {
	r._orderNo = ""
	r._feature = ""
	r._serviceCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.locallife.communitylife.verifyservicecode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoCntecLocallifeCommunitylifeVerifyservicecodeRequest()
	},
}

// GetCainiaoCntecLocallifeCommunitylifeVerifyservicecodeRequest 从 sync.Pool 获取 CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest
func GetCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest() *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest {
	return poolCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest.Get().(*CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest)
}

// ReleaseCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest 将 CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest 放入 sync.Pool
func ReleaseCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest(v *CainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest) {
	v.Reset()
	poolCainiaoCntecLocallifeCommunitylifeVerifyservicecodeAPIRequest.Put(v)
}
