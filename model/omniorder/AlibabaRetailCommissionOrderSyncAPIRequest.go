package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailCommissionOrderSyncAPIRequest 分佣数据传输 API请求
// alibaba.retail.commission.order.sync
//
// 同步分佣结果
type AlibabaRetailCommissionOrderSyncAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *UniverseOrderVo
}

// NewAlibabaRetailCommissionOrderSyncRequest 初始化AlibabaRetailCommissionOrderSyncAPIRequest对象
func NewAlibabaRetailCommissionOrderSyncRequest() *AlibabaRetailCommissionOrderSyncAPIRequest {
	return &AlibabaRetailCommissionOrderSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailCommissionOrderSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailCommissionOrderSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailCommissionOrderSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabaRetailCommissionOrderSyncAPIRequest) SetParam0(_param0 *UniverseOrderVo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailCommissionOrderSyncAPIRequest) GetParam0() *UniverseOrderVo {
	return r._param0
}
