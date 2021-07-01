package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRechargeUndedutUpdateAPIRequest
储值消费退款(逆向) API请求
alibaba.alsc.crm.recharge.undedut.update

新增储值消费退款接口 */
type AlibabaAlscCrmRechargeUndedutUpdateAPIRequest struct {
	model.Params
	// 入参
	_paramUndedutOpenReq *UndedutOpenReq
}

// New
