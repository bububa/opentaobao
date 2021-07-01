package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeAccountflowsGetAPIRequest
分页查询储值流水 API请求
alibaba.alsc.crm.recharge.accountflows.get

增加分页查询储值流水接口 */
type AlibabaAlscCrmRechargeAccountflowsGetAPIRequest struct {
	model.Params
	// 入参
	_paramPageQueryAccountFlowsOpenReq *PageQueryAccountFlowsOpenReq
}

// New
