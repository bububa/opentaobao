package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSpecsTicketGetAPIRequest 产品规格审核信息获取接口 API请求
// tmall.product.specs.ticket.get
//
// 批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等
type TmallProductSpecsTicketGetAPIRequest struct {
	model.Params
	// 产品规格ID，多个用逗号分隔
	_specIds string
}

// NewTmallProductSpecsTicketGetRequest 初始化TmallProductSpecsTicketGetAPIRequest对象
func NewTmallProductSpecsTicketGetRequest() *TmallProductSpecsTicketGetAPIRequest {
	return &TmallProductSpecsTicketGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallProductSpecsTicketGetAPIRequest) Reset() {
	r._specIds = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductSpecsTicketGetAPIRequest) GetApiMethodName() string {
	return "tmall.product.specs.ticket.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductSpecsTicketGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallProductSpecsTicketGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSpecIds is SpecIds Setter
// 产品规格ID，多个用逗号分隔
func (r *TmallProductSpecsTicketGetAPIRequest) SetSpecIds(_specIds string) error {
	r._specIds = _specIds
	r.Set("spec_ids", _specIds)
	return nil
}

// GetSpecIds SpecIds Getter
func (r TmallProductSpecsTicketGetAPIRequest) GetSpecIds() string {
	return r._specIds
}

var poolTmallProductSpecsTicketGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallProductSpecsTicketGetRequest()
	},
}

// GetTmallProductSpecsTicketGetRequest 从 sync.Pool 获取 TmallProductSpecsTicketGetAPIRequest
func GetTmallProductSpecsTicketGetAPIRequest() *TmallProductSpecsTicketGetAPIRequest {
	return poolTmallProductSpecsTicketGetAPIRequest.Get().(*TmallProductSpecsTicketGetAPIRequest)
}

// ReleaseTmallProductSpecsTicketGetAPIRequest 将 TmallProductSpecsTicketGetAPIRequest 放入 sync.Pool
func ReleaseTmallProductSpecsTicketGetAPIRequest(v *TmallProductSpecsTicketGetAPIRequest) {
	v.Reset()
	poolTmallProductSpecsTicketGetAPIRequest.Put(v)
}
