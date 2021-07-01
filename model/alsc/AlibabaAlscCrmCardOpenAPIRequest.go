package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardOpenAPIRequest
标准开卡流程 API请求
alibaba.alsc.crm.card.open

标准开卡流程 */
type AlibabaAlscCrmCardOpenAPIRequest struct {
	model.Params
	// 开卡参数
	_paramOpenCardStandardOpenReq *OpenCardStandardOpenReq
}

// New
