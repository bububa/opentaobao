package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmOpenRechargeOperateAPIRequest
储值操作接口 API请求
alibaba.alsc.crm.open.recharge.operate

储值操作接口 */
type AlibabaAlscCrmOpenRechargeOperateAPIRequest struct {
	model.Params
	// 储值操作参数
	_paramRechargeOperateOpenReq *RechargeOperateOpenReq
}

// New
