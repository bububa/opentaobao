package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateAppversionCreate 创建应用升级任务
// yunos.osupdate.appversion.create
//
// 创建应用升级任务
func YunosOsupdateAppversionCreate(clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionCreateAPIRequest, resp *tvupadmin.YunosOsupdateAppversionCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
