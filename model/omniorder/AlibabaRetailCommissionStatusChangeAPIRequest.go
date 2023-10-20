package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailcommissionstatuschangeAPIRequest 分佣状态变更 API请求
// alibaba.retail.commission.status.change
//
// 分佣系统，分佣状态变更接口
type AlibabaretailcommissionstatuschangeAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *UniverseOrderVo
}

// NewAlibabaretailcommissionstatuschangeRequest 初始化AlibabaretailcommissionstatuschangeAPIRequest对象
func NewAlibabaretailcommissionstatuschangeRequest() *AlibabaretailcommissionstatuschangeAPIRequest {
	return &AlibabaretailcommissionstatuschangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailcommissionstatuschangeAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailcommissionstatuschangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailcommissionstatuschangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabaretailcommissionstatuschangeAPIRequest) SetParam0(_param0 *UniverseOrderVo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaretailcommissionstatuschangeAPIRequest) GetParam0() *UniverseOrderVo {
	return r._param0
}
