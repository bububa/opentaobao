package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGroupsGetAPIRequest
查询卖家的分组 API请求
taobao.crm.groups.get

查询卖家的分组，返回查询到的分组列表，分页返回分组 */
type TaobaoCrmGroupsGetAPIRequest struct {
	model.Params
	// 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条
	_pageSize int64
	// 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1
	_currentPage int64
}

// New
