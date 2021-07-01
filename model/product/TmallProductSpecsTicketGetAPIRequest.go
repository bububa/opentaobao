package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallProductSpecsTicketGetAPIRequest
产品规格审核信息获取接口 API请求
tmall.product.specs.ticket.get

批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等 */
type TmallProductSpecsTicketGetAPIRequest struct {
	model.Params
	// 产品规格ID，多个用逗号分隔
	_specIds string
}

// NewTmallProductSpecsTicketGetRequest 初始化TmallProductSpecsTicketGetAPIRequest对象
func NewTmallProductSpecsTicketGetRequest() *TmallProductSpecsTicketGetAPIRequest {
	return &TmallProductSpecsTicketGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductSpecsTicketGetAPIRequest) GetApiMethodName() string {
	return "tmall.product.specs.ticket.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductSpecsTicketGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SpecIds Setter
// 产品规格ID，多个用逗号分隔
func (r *TmallProductSpecsTicketGetAPIRequest) SetSpecIds(_specIds string) error {
	r._specIds = _specIds
	r.Set("spec_ids", _specIds)
	return nil
}

// Get SpecIds Getter
func (r TmallProductSpecsTicketGetAPIRequest) GetSpecIds() string {
	return r._specIds
}
