package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmMembersIncrementGetAPIRequest
增量获取卖家会员 API请求
taobao.crm.members.increment.get

增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100 */
type TaobaoCrmMembersIncrementGetAPIRequest struct {
	model.Params
	// 会员等级
	_grade int64
	// 卖家修改会员信息的时间起点.
	_startModify string
	// 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间.
	_endModify string
	// 每页显示的会员数，page_size的值不能超过100，最小值要大于1
	_pageSize int64
	// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1
	_currentPage int64
}

// New
