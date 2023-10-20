package qt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqtreportdeleteAPIRequest 质检报告删除接口 API请求
// taobao.qt.report.delete
//
// 删除质检报告
type TaobaoqtreportdeleteAPIRequest struct {
	model.Params
	// 一个质检服务唯一标识质量检验单的编号
	_qtCode string
}

// NewTaobaoqtreportdeleteRequest 初始化TaobaoqtreportdeleteAPIRequest对象
func NewTaobaoqtreportdeleteRequest() *TaobaoqtreportdeleteAPIRequest {
	return &TaobaoqtreportdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqtreportdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qt.report.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqtreportdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqtreportdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQtCode is QtCode Setter
// 一个质检服务唯一标识质量检验单的编号
func (r *TaobaoqtreportdeleteAPIRequest) SetQtCode(_qtCode string) error {
	r._qtCode = _qtCode
	r.Set("qt_code", _qtCode)
	return nil
}

// GetQtCode QtCode Getter
func (r TaobaoqtreportdeleteAPIRequest) GetQtCode() string {
	return r._qtCode
}
