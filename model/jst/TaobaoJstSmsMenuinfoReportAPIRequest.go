package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsMenuinfoReportAPIRequest 聚石塔菜单信息上报 API请求
// taobao.jst.sms.menuinfo.report
//
// 聚石塔菜单信息上报
type TaobaoJstSmsMenuinfoReportAPIRequest struct {
	model.Params
	// 菜单信息上报接口的请求参数
	_menuInfoReportRequest *MenuInfoReportRequest
}

// NewTaobaoJstSmsMenuinfoReportRequest 初始化TaobaoJstSmsMenuinfoReportAPIRequest对象
func NewTaobaoJstSmsMenuinfoReportRequest() *TaobaoJstSmsMenuinfoReportAPIRequest {
	return &TaobaoJstSmsMenuinfoReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMenuinfoReportAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.menuinfo.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMenuinfoReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MenuInfoReportRequest Setter
// 菜单信息上报接口的请求参数
func (r *TaobaoJstSmsMenuinfoReportAPIRequest) SetMenuInfoReportRequest(_menuInfoReportRequest *MenuInfoReportRequest) error {
	r._menuInfoReportRequest = _menuInfoReportRequest
	r.Set("menu_info_report_request", _menuInfoReportRequest)
	return nil
}

// Get MenuInfoReportRequest Getter
func (r TaobaoJstSmsMenuinfoReportAPIRequest) GetMenuInfoReportRequest() *MenuInfoReportRequest {
	return r._menuInfoReportRequest
}
