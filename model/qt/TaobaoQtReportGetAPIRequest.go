package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqtreportgetAPIRequest 查询质检报告 API请求
// taobao.qt.report.get
//
// 质检报告查询
type TaobaoqtreportgetAPIRequest struct {
	model.Params
	// 质检编号
	_qtCode string
}

// NewTaobaoqtreportgetRequest 初始化TaobaoqtreportgetAPIRequest对象
func NewTaobaoqtreportgetRequest() *TaobaoqtreportgetAPIRequest {
	return &TaobaoqtreportgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqtreportgetAPIRequest) GetApiMethodName() string {
	return "taobao.qt.report.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqtreportgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqtreportgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQtCode is QtCode Setter
// 质检编号
func (r *TaobaoqtreportgetAPIRequest) SetQtCode(_qtCode string) error {
	r._qtCode = _qtCode
	r.Set("qt_code", _qtCode)
	return nil
}

// GetQtCode QtCode Getter
func (r TaobaoqtreportgetAPIRequest) GetQtCode() string {
	return r._qtCode
}
