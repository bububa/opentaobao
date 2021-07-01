package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServiceSettleadjustmentGetAPIRequest
查询结算调整单单条记录 API请求
tmall.service.settleadjustment.get

提供给服务商通过结算调整单id获取结算调整单信息 */
type TmallServiceSettleadjustmentGetAPIRequest struct {
	model.Params
	// 结算调整单ID
	_id int64
}

// New
