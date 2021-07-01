package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMaitixOrderCancelAPIRequest
大麦-库存释放 API请求
alibaba.damai.maitix.order.cancel

库存释放 */
type AlibabaDamaiMaitixOrderCancelAPIRequest struct {
	model.Params
	// 库存释放入参
	_param *MoaUnlockTicketParam
}

// NewAlibabaDamaiMaitixOrderCancelRequest 初始化AlibabaDamaiMaitixOrderCancelAPIRequest对象
func NewAlibabaDamaiMaitixOrderCancelRequest() *AlibabaDamaiMaitixOrderCancelAPIRequest {
	return &AlibabaDamaiMaitixOrderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 库存释放入参
func (r *AlibabaDamaiMaitixOrderCancelAPIRequest) SetParam(_param *MoaUnlockTicketParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaDamaiMaitixOrderCancelAPIRequest) GetParam() *MoaUnlockTicketParam {
	return r._param
}
