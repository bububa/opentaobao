package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

/* AlibabaJymRequirementOrdersOperationNotify
交易猫需求订单操作接口
alibaba.jym.requirement.orders.operation.notify

交易猫需求订单操作接口 */
func AlibabaJymRequirementOrdersOperationNotify(clt *core.SDKClient, req *jym.AlibabaJymRequirementOrdersOperationNotifyAPIRequest, session string) (*jym.AlibabaJymRequirementOrdersOperationNotifyAPIResponse, error) {
	var resp jym.AlibabaJymRequirementOrdersOperationNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
