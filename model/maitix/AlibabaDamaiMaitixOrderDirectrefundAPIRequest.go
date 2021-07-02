package maitix

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderDirectrefundAPIRequest 大麦-直接退票 API请求
// alibaba.damai.maitix.order.directrefund
//
// 大麦-退票
type AlibabaDamaiMaitixOrderDirectrefundAPIRequest struct {
	model.Params
	// 退票入参
	_param *MoaRefundAuditParam
}

// NewAlibabaDamaiMaitixOrderDirectrefundRequest 初始化AlibabaDamaiMaitixOrderDirectrefundAPIRequest对象
func NewAlibabaDamaiMaitixOrderDirectrefundRequest() *AlibabaDamaiMaitixOrderDirectrefundAPIRequest {
	return &AlibabaDamaiMaitixOrderDirectrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixOrderDirectrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.order.directrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixOrderDirectrefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 退票入参
func (r *AlibabaDamaiMaitixOrderDirectrefundAPIRequest) SetParam(_param *MoaRefundAuditParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlibabaDamaiMaitixOrderDirectrefundAPIRequest) GetParam() *MoaRefundAuditParam {
	return r._param
}
