package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionStatusChangeAPIRequest 分佣状态变更 API请求
// alibaba.retail.commission.status.change
//
// 分佣系统，分佣状态变更接口
type AlibabaRetailCommissionStatusChangeAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *UniverseOrderVo
}

// NewAlibabaRetailCommissionStatusChangeRequest 初始化AlibabaRetailCommissionStatusChangeAPIRequest对象
func NewAlibabaRetailCommissionStatusChangeRequest() *AlibabaRetailCommissionStatusChangeAPIRequest {
	return &AlibabaRetailCommissionStatusChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailCommissionStatusChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailCommissionStatusChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 请求参数
func (r *AlibabaRetailCommissionStatusChangeAPIRequest) SetParam0(_param0 *UniverseOrderVo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaRetailCommissionStatusChangeAPIRequest) GetParam0() *UniverseOrderVo {
	return r._param0
}
