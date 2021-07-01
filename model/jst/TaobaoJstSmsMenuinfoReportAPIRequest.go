package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsMenuinfoReportAPIRequest
聚石塔菜单信息上报 API请求
taobao.jst.sms.menuinfo.report

聚石塔菜单信息上报 */
type TaobaoJstSmsMenuinfoReportAPIRequest struct {
	model.Params
	// 菜单信息上报接口的请求参数
	_menuInfoReportRequest *MenuInfoReportRequest
}

// New
