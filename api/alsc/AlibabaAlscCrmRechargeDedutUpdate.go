package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmRechargeDedutUpdate
储值消费
alibaba.alsc.crm.recharge.dedut.update

增加储值消费接口 */
func AlibabaAlscCrmRechargeDedutUpdate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeDedutUpdateAPIRequest, session string) (*alsc.AlibabaAlscCrmRechargeDedutUpdateAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmRechargeDedutUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
