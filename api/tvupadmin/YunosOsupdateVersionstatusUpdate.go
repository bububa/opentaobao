package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateVersionstatusUpdate 更新应用升级状态
// yunos.osupdate.versionstatus.update
//
// 更新应用升级状态
func YunosOsupdateVersionstatusUpdate(clt *core.SDKClient, req *tvupadmin.YunosOsupdateVersionstatusUpdateAPIRequest, resp *tvupadmin.YunosOsupdateVersionstatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
