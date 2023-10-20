package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanrawpogapreturnAPIRequest 二级物料-LT内的POGAP数据回传 API请求
// alibaba.tmallgenie.scp.plan.rawpo.gap.return
//
// 二级物料-LT内的POGAP数据回传
type AlibabatmallgeniescpplanrawpogapreturnAPIRequest struct {
	model.Params
	// 请求对象
	_rawPogapRequest *RawPurchaseOrderGapRequest
}

// NewAlibabatmallgeniescpplanrawpogapreturnRequest 初始化AlibabatmallgeniescpplanrawpogapreturnAPIRequest对象
func NewAlibabatmallgeniescpplanrawpogapreturnRequest() *AlibabatmallgeniescpplanrawpogapreturnAPIRequest {
	return &AlibabatmallgeniescpplanrawpogapreturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanrawpogapreturnAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.rawpo.gap.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanrawpogapreturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanrawpogapreturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRawPogapRequest is RawPogapRequest Setter
// 请求对象
func (r *AlibabatmallgeniescpplanrawpogapreturnAPIRequest) SetRawPogapRequest(_rawPogapRequest *RawPurchaseOrderGapRequest) error {
	r._rawPogapRequest = _rawPogapRequest
	r.Set("raw_pogap_request", _rawPogapRequest)
	return nil
}

// GetRawPogapRequest RawPogapRequest Getter
func (r AlibabatmallgeniescpplanrawpogapreturnAPIRequest) GetRawPogapRequest() *RawPurchaseOrderGapRequest {
	return r._rawPogapRequest
}
