package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoRequisitionsGetAPIRequest
合作申请查询 API请求
taobao.fenxiao.requisitions.get

合作申请查询 */
type TaobaoFenxiaoRequisitionsGetAPIRequest struct {
	model.Params
	// 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期）
	_status int64
	// 申请开始时间yyyy-MM-dd
	_applyStart string
	// 申请结束时间yyyy-MM-dd
	_applyEnd string
	// 页码（大于0的整数，默认1）
	_pageNo int64
	// 每页记录数（默认20，最大50）
	_pageSize int64
}

// New
