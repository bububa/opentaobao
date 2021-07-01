package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest
储值流水详细 API请求
alibaba.alsc.crm.recharge.account.flowdetail.get

查询储值流水详细接口 */
type AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryRechargeAccountFlowOpenReq *QueryRechargeAccountFlowOpenReq
}

// New
