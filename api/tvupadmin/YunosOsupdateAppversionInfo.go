package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateAppversionInfo 获取应用升级详情
// yunos.osupdate.appversion.info
//
// 获取应用升级详情
func YunosOsupdateAppversionInfo(clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionInfoAPIRequest, resp *tvupadmin.YunosOsupdateAppversionInfoAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
