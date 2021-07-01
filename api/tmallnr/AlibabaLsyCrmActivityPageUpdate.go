package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

/* AlibabaLsyCrmActivityPageUpdate
ISV活动页面创建修改
alibaba.lsy.crm.activity.page.update

ISV活动页面创建修改 */
func AlibabaLsyCrmActivityPageUpdate(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityPageUpdateAPIRequest, session string) (*tmallnr.AlibabaLsyCrmActivityPageUpdateAPIResponse, error) {
	var resp tmallnr.AlibabaLsyCrmActivityPageUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
