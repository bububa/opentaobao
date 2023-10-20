package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductspecsticketgetAPIRequest 产品规格审核信息获取接口 API请求
// tmall.product.specs.ticket.get
//
// 批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等
type TmallproductspecsticketgetAPIRequest struct {
	model.Params
	// 产品规格ID，多个用逗号分隔
	_specIds string
}

// NewTmallproductspecsticketgetRequest 初始化TmallproductspecsticketgetAPIRequest对象
func NewTmallproductspecsticketgetRequest() *TmallproductspecsticketgetAPIRequest {
	return &TmallproductspecsticketgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductspecsticketgetAPIRequest) GetApiMethodName() string {
	return "tmall.product.specs.ticket.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductspecsticketgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductspecsticketgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpecIds is SpecIds Setter
// 产品规格ID，多个用逗号分隔
func (r *TmallproductspecsticketgetAPIRequest) SetSpecIds(_specIds string) error {
	r._specIds = _specIds
	r.Set("spec_ids", _specIds)
	return nil
}

// GetSpecIds SpecIds Getter
func (r TmallproductspecsticketgetAPIRequest) GetSpecIds() string {
	return r._specIds
}
