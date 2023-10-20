package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateOsfotaAdd 添加系统升级任务
// yunos.osupdate.osfota.add
//
// 添加osupdate系统升级任务
func YunosOsupdateOsfotaAdd(clt *core.SDKClient, req *tvupadmin.YunosOsupdateOsfotaAddAPIRequest, resp *tvupadmin.YunosOsupdateOsfotaAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
