package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoCooperationGetAPIRequest
供应商或分销商获取合作关系信息 API请求
taobao.fenxiao.cooperation.get

获取供应商的合作关系信息 */
type TaobaoFenxiaoCooperationGetAPIRequest struct {
	model.Params
	// 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止)
	_status string
	// 合作开始时间yyyy-MM-dd HH:mm:ss
	_startDate string
	// 合作结束时间yyyy-MM-dd HH:mm:ss
	_endDate string
	// 分销方式：AGENT(代销) 、DEALER（经销）
	_tradeType string
	// 页码（大于0的整数，默认1）
	_pageNo int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
	// 渠道code
	_channelCode string
	// 1是供应商，2 是分销商
	_roleType int64
}

// New
