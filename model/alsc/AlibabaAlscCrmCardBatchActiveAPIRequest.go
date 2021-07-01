package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardBatchActiveAPIRequest
批量激活卡 API请求
alibaba.alsc.crm.card.batch.active

批量激活卡 */
type AlibabaAlscCrmCardBatchActiveAPIRequest struct {
	model.Params
	// 请求对象
	_paramBatchActiveCardOpenReq *BatchActiveCardOpenReq
}

// New
