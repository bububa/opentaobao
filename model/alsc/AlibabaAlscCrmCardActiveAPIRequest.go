package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardActiveAPIRequest
标准激活卡 API请求
alibaba.alsc.crm.card.active

激活卡 */
type AlibabaAlscCrmCardActiveAPIRequest struct {
	model.Params
	// 请求参数
	_paramActiveCardOpenReq *ActiveCardOpenReq
}

// New
