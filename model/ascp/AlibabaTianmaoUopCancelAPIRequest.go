package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianmaouopcancelAPIRequest 阿里巴巴.天猫. 履约订单. 取消 API请求
// alibaba.tianmao.uop.cancel
//
// 阿里巴巴.天猫. 履约订单. 取消
type AlibabatianmaouopcancelAPIRequest struct {
	model.Params
	// 取消接口入参
	_hiErpCloseDto *HiErpCloseDto
}

// NewAlibabatianmaouopcancelRequest 初始化AlibabatianmaouopcancelAPIRequest对象
func NewAlibabatianmaouopcancelRequest() *AlibabatianmaouopcancelAPIRequest {
	return &AlibabatianmaouopcancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianmaouopcancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tianmao.uop.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianmaouopcancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianmaouopcancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHiErpCloseDto is HiErpCloseDto Setter
// 取消接口入参
func (r *AlibabatianmaouopcancelAPIRequest) SetHiErpCloseDto(_hiErpCloseDto *HiErpCloseDto) error {
	r._hiErpCloseDto = _hiErpCloseDto
	r.Set("hi_erp_close_dto", _hiErpCloseDto)
	return nil
}

// GetHiErpCloseDto HiErpCloseDto Getter
func (r AlibabatianmaouopcancelAPIRequest) GetHiErpCloseDto() *HiErpCloseDto {
	return r._hiErpCloseDto
}
