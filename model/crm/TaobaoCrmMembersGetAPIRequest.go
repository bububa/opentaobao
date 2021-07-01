package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmMembersGetAPIRequest
获取卖家的会员（基本查询） API请求
taobao.crm.members.get

查询卖家的会员，进行基本的查询，返回符合条件的会员列表 */
type TaobaoCrmMembersGetAPIRequest struct {
	model.Params
	// 买家的昵称
	_buyerNick string
	// 会员等级,如果不传入值则默认为全部等级。
	_grade int64
	// 最小交易额,单位为元
	_minTradeAmount float64
	// 最大交易额，单位为元
	_maxTradeAmount float64
	// 最小交易量
	_minTradeCount int64
	// 最大交易量
	_maxTradeCount int64
	// 最早上次交易时间
	_minLastTradeTime string
	// 最迟上次交易时间
	_maxLastTradeTime string
	// 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1，
	_pageSize int64
	// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为1000
	_currentPage int64
}

// New
