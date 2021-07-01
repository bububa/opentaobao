package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeDedutUpdateAPIRequest
储值消费 API请求
alibaba.alsc.crm.recharge.dedut.update

增加储值消费接口 */
type AlibabaAlscCrmRechargeDedutUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramDedutOpenReq *DedutOpenReq
}

// New
