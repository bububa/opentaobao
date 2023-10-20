package cainiaocntec

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest 订单状态推送 API请求
// cainiao.cntec.locallife.communitylife.syncorderstatus
//
// 订单状态推送
type CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest struct {
	model.Params
	// 订单id
	_orderNo string
	// 订单状态描述
	_statusDesc string
	// 扩展信息json字段
	_feature string
	// 订单状态码
	_statusCode string
}

// NewCainiaocnteclocallifecommunitylifesyncorderstatusRequest 初始化CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest对象
func NewCainiaocnteclocallifecommunitylifesyncorderstatusRequest() *CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest {
	return &CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.locallife.communitylife.syncorderstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 订单id
func (r *CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetStatusDesc is StatusDesc Setter
// 订单状态描述
func (r *CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) SetStatusDesc(_statusDesc string) error {
	r._statusDesc = _statusDesc
	r.Set("status_desc", _statusDesc)
	return nil
}

// GetStatusDesc StatusDesc Getter
func (r CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) GetStatusDesc() string {
	return r._statusDesc
}

// SetFeature is Feature Setter
// 扩展信息json字段
func (r *CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) GetFeature() string {
	return r._feature
}

// SetStatusCode is StatusCode Setter
// 订单状态码
func (r *CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) SetStatusCode(_statusCode string) error {
	r._statusCode = _statusCode
	r.Set("status_code", _statusCode)
	return nil
}

// GetStatusCode StatusCode Getter
func (r CainiaocnteclocallifecommunitylifesyncorderstatusAPIRequest) GetStatusCode() string {
	return r._statusCode
}
