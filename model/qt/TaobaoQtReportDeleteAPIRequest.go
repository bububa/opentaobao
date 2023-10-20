package qt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQtReportDeleteAPIRequest 质检报告删除接口 API请求
// taobao.qt.report.delete
//
// 删除质检报告
type TaobaoQtReportDeleteAPIRequest struct {
	model.Params
	// 一个质检服务唯一标识质量检验单的编号
	_qtCode string
}

// NewTaobaoQtReportDeleteRequest 初始化TaobaoQtReportDeleteAPIRequest对象
func NewTaobaoQtReportDeleteRequest() *TaobaoQtReportDeleteAPIRequest {
	return &TaobaoQtReportDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQtReportDeleteAPIRequest) Reset() {
	r._qtCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQtReportDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qt.report.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQtReportDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQtReportDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQtCode is QtCode Setter
// 一个质检服务唯一标识质量检验单的编号
func (r *TaobaoQtReportDeleteAPIRequest) SetQtCode(_qtCode string) error {
	r._qtCode = _qtCode
	r.Set("qt_code", _qtCode)
	return nil
}

// GetQtCode QtCode Getter
func (r TaobaoQtReportDeleteAPIRequest) GetQtCode() string {
	return r._qtCode
}

var poolTaobaoQtReportDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQtReportDeleteRequest()
	},
}

// GetTaobaoQtReportDeleteRequest 从 sync.Pool 获取 TaobaoQtReportDeleteAPIRequest
func GetTaobaoQtReportDeleteAPIRequest() *TaobaoQtReportDeleteAPIRequest {
	return poolTaobaoQtReportDeleteAPIRequest.Get().(*TaobaoQtReportDeleteAPIRequest)
}

// ReleaseTaobaoQtReportDeleteAPIRequest 将 TaobaoQtReportDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoQtReportDeleteAPIRequest(v *TaobaoQtReportDeleteAPIRequest) {
	v.Reset()
	poolTaobaoQtReportDeleteAPIRequest.Put(v)
}
