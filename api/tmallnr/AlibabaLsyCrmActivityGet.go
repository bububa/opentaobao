package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivityGet 私域导购查询活动详情
// alibaba.lsy.crm.activity.get
//
// 私域导购查询活动详情
func AlibabaLsyCrmActivityGet(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityGetAPIRequest, resp *tmallnr.AlibabaLsyCrmActivityGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
