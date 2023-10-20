package omniorder

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailCommissionStatusChangeAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailCommissionStatusChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.commission.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailCommissionStatusChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailCommissionStatusChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabaRetailCommissionStatusChangeAPIRequest) SetParam0(_param0 *UniverseOrderVo) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaRetailCommissionStatusChangeAPIRequest) GetParam0() *UniverseOrderVo {
	return r._param0
}

var poolAlibabaRetailCommissionStatusChangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailCommissionStatusChangeRequest()
	},
}

// GetAlibabaRetailCommissionStatusChangeRequest 从 sync.Pool 获取 AlibabaRetailCommissionStatusChangeAPIRequest
func GetAlibabaRetailCommissionStatusChangeAPIRequest() *AlibabaRetailCommissionStatusChangeAPIRequest {
	return poolAlibabaRetailCommissionStatusChangeAPIRequest.Get().(*AlibabaRetailCommissionStatusChangeAPIRequest)
}

// ReleaseAlibabaRetailCommissionStatusChangeAPIRequest 将 AlibabaRetailCommissionStatusChangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailCommissionStatusChangeAPIRequest(v *AlibabaRetailCommissionStatusChangeAPIRequest) {
	v.Reset()
	poolAlibabaRetailCommissionStatusChangeAPIRequest.Put(v)
}
