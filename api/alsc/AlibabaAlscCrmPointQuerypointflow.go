package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmPointQuerypointflow
分页查询积分流水
alibaba.alsc.crm.point.querypointflow

分页查询积分流水 */
func AlibabaAlscCrmPointQuerypointflow(clt *core.SDKClient, req *alsc.AlibabaAlscCrmPointQuerypointflowAPIRequest, session string) (*alsc.AlibabaAlscCrmPointQuerypointflowAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmPointQuerypointflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
