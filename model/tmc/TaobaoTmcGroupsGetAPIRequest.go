package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcGroupsGetAPIRequest
获取自定义用户分组列表 API请求
taobao.tmc.groups.get

获取自定义用户分组列表 */
type TaobaoTmcGroupsGetAPIRequest struct {
	model.Params
	// 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。
	_groupNames []string
	// 页码
	_pageNo int64
	// 每页返回多少个分组
	_pageSize int64
}

// New
