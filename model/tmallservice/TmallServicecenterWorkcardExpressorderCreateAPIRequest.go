package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterWorkcardExpressorderCreateAPIRequest
物流订单创建API API请求
tmall.servicecenter.workcard.expressorder.create

天猫服务寄送类业务，服务商履约完成后寄回操作时，提供的物流寄件单创建API */
type TmallServicecenterWorkcardExpressorderCreateAPIRequest struct {
	model.Params
	// 物流单关联的工单List
	_workcardIdList []int64
	// 真实履约服务商Nick（非ERP系统不要填写）
	_realTpNick string
	// erp外部物流订单号
	_externalLogisticsOrderId string
}

// New
