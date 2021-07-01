package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardQryphysicalAPIRequest
查询物理卡 API请求
alibaba.alsc.crm.card.qryphysical

查询物理卡 */
type AlibabaAlscCrmCardQryphysicalAPIRequest struct {
	model.Params
	// 入参
	_paramQueryPhyCardOpenReq *QueryPhyCardOpenReq
}

// New
