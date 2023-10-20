package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateOsfotaPublish 系统升级发布
// yunos.osupdate.osfota.publish
//
// 发布osupdate系统升级任务
func YunosOsupdateOsfotaPublish(clt *core.SDKClient, req *tvupadmin.YunosOsupdateOsfotaPublishAPIRequest, resp *tvupadmin.YunosOsupdateOsfotaPublishAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
