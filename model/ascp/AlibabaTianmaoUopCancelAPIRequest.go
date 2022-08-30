package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianmaoUopCancelAPIRequest 阿里巴巴.天猫. 履约订单. 取消 API请求
// alibaba.tianmao.uop.cancel
//
// 阿里巴巴.天猫. 履约订单. 取消
type AlibabaTianmaoUopCancelAPIRequest struct {
	model.Params
	// 取消接口入参
	_hiErpCloseDto *HiErpCloseDto
}

// NewAlibabaTianmaoUopCancelRequest 初始化AlibabaTianmaoUopCancelAPIRequest对象
func NewAlibabaTianmaoUopCancelRequest() *AlibabaTianmaoUopCancelAPIRequest {
	return &AlibabaTianmaoUopCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianmaoUopCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.uop.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianmaoUopCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetHiErpCloseDto is HiErpCloseDto Setter
// 取消接口入参
func (r *AlibabaTianmaoUopCancelAPIRequest) SetHiErpCloseDto(_hiErpCloseDto *HiErpCloseDto) error {
	r._hiErpCloseDto = _hiErpCloseDto
	r.Set("hi_erp_close_dto", _hiErpCloseDto)
	return nil
}

// GetHiErpCloseDto HiErpCloseDto Getter
func (r AlibabaTianmaoUopCancelAPIRequest) GetHiErpCloseDto() *HiErpCloseDto {
	return r._hiErpCloseDto
}
