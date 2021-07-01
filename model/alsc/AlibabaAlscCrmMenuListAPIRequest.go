package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmMenuListAPIRequest
获取特价菜单 API请求
alibaba.alsc.crm.menu.list

获取特价菜单 */
type AlibabaAlscCrmMenuListAPIRequest struct {
	model.Params
	// 获取特价菜单请求参数
	_menuOpenReq *MenuOpenReq
}

// New
