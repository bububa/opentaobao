package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoNewretailDivisionRecordListGetAPIRequest
导购分佣明细列表 API请求
taobao.newretail.division.record.list.get

提供分页查询导购分佣明细的能力 */
type TaobaoNewretailDivisionRecordListGetAPIRequest struct {
	model.Params
	// 入参
	_param *TopDivisionRecordReqDto
}

// New
