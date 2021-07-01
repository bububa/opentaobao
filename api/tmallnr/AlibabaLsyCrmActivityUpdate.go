package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

/* AlibabaLsyCrmActivityUpdate
ISV活动修改
alibaba.lsy.crm.activity.update

ISV活动修改 */
func AlibabaLsyCrmActivityUpdate(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityUpdateAPIRequest, session string) (*tmallnr.AlibabaLsyCrmActivityUpdateAPIResponse, error) {
	var resp tmallnr.AlibabaLsyCrmActivityUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
