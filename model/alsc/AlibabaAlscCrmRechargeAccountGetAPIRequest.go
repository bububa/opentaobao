package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeAccountGetAPIRequest
查询储值账户信息 API请求
alibaba.alsc.crm.recharge.account.get

查询储值账户信息接口 */
type AlibabaAlscCrmRechargeAccountGetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryRechargeAccountOpenReq *QueryRechargeAccountOpenReq
}

// New
