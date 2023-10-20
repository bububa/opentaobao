package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailcommissionordersyncAPIRequest 分佣数据传输 API请求
// alibaba.retail.commission.order.sync
//
// 同步分佣结果
type AlibabaretailcommissionordersyncAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *UniverseOrderVo
}

// NewAlibabaretailcommissionordersyncRequest 初始化AlibabaretailcommissionordersyncAPIRequest对象
func NewAlibabaretailcommissionordersyncRequest() *AlibabaretailcommissionordersyncAPIRequest {
	return &AlibabaretailcommissionordersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaretailcommissionordersyncAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaretailcommissionordersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaretailcommissionordersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabaretailcommissionordersyncAPIRequest) SetParam0(_param0 *UniverseOrderVo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaretailcommissionordersyncAPIRequest) GetParam0() *UniverseOrderVo {
	return r._param0
}
