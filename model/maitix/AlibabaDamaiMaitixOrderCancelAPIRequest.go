package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixordercancelAPIRequest 大麦-库存释放 API请求
// alibaba.damai.maitix.order.cancel
//
// 库存释放
type AlibabadamaimaitixordercancelAPIRequest struct {
	model.Params
	// 库存释放入参
	_param *MoaUnlockTicketParam
}

// NewAlibabadamaimaitixordercancelRequest 初始化AlibabadamaimaitixordercancelAPIRequest对象
func NewAlibabadamaimaitixordercancelRequest() *AlibabadamaimaitixordercancelAPIRequest {
	return &AlibabadamaimaitixordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimaitixordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimaitixordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimaitixordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 库存释放入参
func (r *AlibabadamaimaitixordercancelAPIRequest) SetParam(_param *MoaUnlockTicketParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabadamaimaitixordercancelAPIRequest) GetParam() *MoaUnlockTicketParam {
	return r._param
}
