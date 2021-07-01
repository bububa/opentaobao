package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardQryAPIRequest
查询卡实例 API请求
alibaba.alsc.crm.card.qry

查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询） */
type AlibabaAlscCrmCardQryAPIRequest struct {
	model.Params
	// 请求对象
	_paramQueryCardOpenReq *QueryCardOpenReq
}

// New
