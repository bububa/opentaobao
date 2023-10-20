package cainiaocntec

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest 订单状态推送 API请求
// cainiao.cntec.locallife.communitylife.syncorderstatus
//
// 订单状态推送
type CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest struct {
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

// NewCainiaoCntecLocallifeCommunitylifeSyncorderstatusRequest 初始化CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest对象
func NewCainiaoCntecLocallifeCommunitylifeSyncorderstatusRequest() *CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest {
	return &CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) GetApiMethodName() string {
	return "cainiao.cntec.locallife.communitylife.syncorderstatus"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 订单id
func (r *CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetStatusDesc is StatusDesc Setter
// 订单状态描述
func (r *CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) SetStatusDesc(_statusDesc string) error {
	r._statusDesc = _statusDesc
	r.Set("status_desc", _statusDesc)
	return nil
}

// GetStatusDesc StatusDesc Getter
func (r CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) GetStatusDesc() string {
	return r._statusDesc
}

// SetFeature is Feature Setter
// 扩展信息json字段
func (r *CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// GetFeature Feature Getter
func (r CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) GetFeature() string {
	return r._feature
}

// SetStatusCode is StatusCode Setter
// 订单状态码
func (r *CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) SetStatusCode(_statusCode string) error {
	r._statusCode = _statusCode
	r.Set("status_code", _statusCode)
	return nil
}

// GetStatusCode StatusCode Getter
func (r CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest) GetStatusCode() string {
	return r._statusCode
}
