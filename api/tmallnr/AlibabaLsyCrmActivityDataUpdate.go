package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// AlibabaLsyCrmActivityDataUpdate 私域导购数据回流接口
// alibaba.lsy.crm.activity.data.update
//
// 私域导购数据回流接口
func AlibabaLsyCrmActivityDataUpdate(clt *core.SDKClient, req *tmallnr.AlibabaLsyCrmActivityDataUpdateAPIRequest, resp *tmallnr.AlibabaLsyCrmActivityDataUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
