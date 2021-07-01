package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmMembersSearchAPIRequest
获取卖家会员（高级查询） API请求
taobao.crm.members.search

会员列表的高级查询，接口返回符合条件的会员列表.<br><br/>注：建议获取09年以后的数据，09年之前的数据不是很完整 */
type TaobaoCrmMembersSearchAPIRequest struct {
	model.Params
	// 买家昵称
	_buyerNick string
	// 会员等级
	_grade int64
	// 最小交易额，单位为元
	_minTradeAmount float64
	// 最大交易额，单位为元
	_maxTradeAmount float64
	// 最小交易量
	_minTradeCount int64
	// 最大交易量
	_maxTradeCount int64
	// 最早上次交易时间（订单创建时间）
	_minLastTradeTime string
	// 最迟上次交易时间
	_maxLastTradeTime string
	// 关系来源，1交易成功，2未成交，3卖家手动吸纳
	_relationSource int64
	// 分组id
	_groupId int64
	// 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1
	_pageSize int64
	// 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页
	_currentPage int64
}

// New
