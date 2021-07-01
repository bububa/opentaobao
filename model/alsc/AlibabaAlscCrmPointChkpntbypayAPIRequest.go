package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPointChkpntbypayAPIRequest
校验支付链路中的积分抵扣是否合法 API请求
alibaba.alsc.crm.point.chkpntbypay

校验支付链路中的积分抵扣是否合法 */
type AlibabaAlscCrmPointChkpntbypayAPIRequest struct {
	model.Params
	// 入参
	_paramConsumePointByPayOpenReq *ConsumePointByPayOpenReq
}

// New
