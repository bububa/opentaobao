package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthreservedentalstoresanditemsAPIRequest 查询商户门店，商品列表 API请求
// alibaba.alihealth.reserve.dental.storesanditems
//
// 查询商户门店，商品列表
type AlibabaalihealthreservedentalstoresanditemsAPIRequest struct {
	model.Params
	// 页码，每页100个门店，超过100个门店分页请求
	_pageNo int64
}

// NewAlibabaalihealthreservedentalstoresanditemsRequest 初始化AlibabaalihealthreservedentalstoresanditemsAPIRequest对象
func NewAlibabaalihealthreservedentalstoresanditemsRequest() *AlibabaalihealthreservedentalstoresanditemsAPIRequest {
	return &AlibabaalihealthreservedentalstoresanditemsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthreservedentalstoresanditemsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.reserve.dental.storesanditems"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthreservedentalstoresanditemsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthreservedentalstoresanditemsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPageNo is PageNo Setter
// 页码，每页100个门店，超过100个门店分页请求
func (r *AlibabaalihealthreservedentalstoresanditemsAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabaalihealthreservedentalstoresanditemsAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
