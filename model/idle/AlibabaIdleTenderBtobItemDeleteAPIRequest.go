package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderbtobitemdeleteAPIRequest 暗拍b2b商品下架/删除 API请求
// alibaba.idle.tender.btob.item.delete
//
// 暗拍b2b商品下架/删除
type AlibabaidletenderbtobitemdeleteAPIRequest struct {
	model.Params
	// 参数0
	_param0 *TenderItemDeleteCmd
}

// NewAlibabaidletenderbtobitemdeleteRequest 初始化AlibabaidletenderbtobitemdeleteAPIRequest对象
func NewAlibabaidletenderbtobitemdeleteRequest() *AlibabaidletenderbtobitemdeleteAPIRequest {
	return &AlibabaidletenderbtobitemdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidletenderbtobitemdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.btob.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidletenderbtobitemdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidletenderbtobitemdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数0
func (r *AlibabaidletenderbtobitemdeleteAPIRequest) SetParam0(_param0 *TenderItemDeleteCmd) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaidletenderbtobitemdeleteAPIRequest) GetParam0() *TenderItemDeleteCmd {
	return r._param0
}
