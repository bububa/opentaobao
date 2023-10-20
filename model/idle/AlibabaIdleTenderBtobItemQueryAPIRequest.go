package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderbtobitemqueryAPIRequest 暗拍b2b商品查询 API请求
// alibaba.idle.tender.btob.item.query
//
// 暗拍b2b商品查询
type AlibabaidletenderbtobitemqueryAPIRequest struct {
	model.Params
	// 参数
	_param0 *TenderItemListQry
}

// NewAlibabaidletenderbtobitemqueryRequest 初始化AlibabaidletenderbtobitemqueryAPIRequest对象
func NewAlibabaidletenderbtobitemqueryRequest() *AlibabaidletenderbtobitemqueryAPIRequest {
	return &AlibabaidletenderbtobitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletenderbtobitemqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.btob.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletenderbtobitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletenderbtobitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *AlibabaidletenderbtobitemqueryAPIRequest) SetParam0(_param0 *TenderItemListQry) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaidletenderbtobitemqueryAPIRequest) GetParam0() *TenderItemListQry {
	return r._param0
}
