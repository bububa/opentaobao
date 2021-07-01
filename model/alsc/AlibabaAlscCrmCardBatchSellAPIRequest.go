package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardBatchSellAPIRequest
批量开卡（售卡） API请求
alibaba.alsc.crm.card.batch.sell

批量开卡（售卡） */
type AlibabaAlscCrmCardBatchSellAPIRequest struct {
	model.Params
	// 请求对象
	_paramBatchOpenCardOpenReq *BatchOpenCardOpenReq
}

// New
