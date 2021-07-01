package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVasSubscSearchAPIRequest
订购记录导出 API请求
taobao.vas.subsc.search

用于ISV查询自己名下的应用及收费项目的订购记录 */
type TaobaoVasSubscSearchAPIRequest struct {
	model.Params
	// 一页包含的记录数
	_pageSize int64
	// 页码
	_pageNo int64
	// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
	_articleCode string
	// 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
	_itemCode string
	// 到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据）
	_startDeadline string
	// 到期时间结束值
	_endDeadline string
	// 订购记录状态，1=有效 2=过期 空=全部
	_status int64
	// 是否自动续费，true=自动续费 false=非自动续费 空=全部
	_autosub bool
	// 是否到期提醒，true=到期提醒 false=非到期提醒 空=全部
	_expireNotice bool
	// 淘宝会员名
	_nick string
}

// New
