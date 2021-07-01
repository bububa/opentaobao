package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmCardBindcustomer
卡号绑定顾客
alibaba.alsc.crm.card.bindcustomer

为卡号绑定顾客 */
func AlibabaAlscCrmCardBindcustomer(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardBindcustomerAPIRequest, session string) (*alsc.AlibabaAlscCrmCardBindcustomerAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmCardBindcustomerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
