package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest
储值核销预先校验 API请求
alibaba.alsc.crm.recharge.dedutprecheck.get

储值核销预先校验接口 */
type AlibabaAlscCrmRechargeDedutprecheckGetAPIRequest struct {
	model.Params
	// 入参
	_paramDeductPreCheckOpenReq *DeductPreCheckOpenReq
}

// New
