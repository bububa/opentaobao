package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateDeviceserviceSearchmodels 根据关键词检索设备型号
// yunos.osupdate.deviceservice.searchmodels
//
// 根据关键词检索设备型号
func YunosOsupdateDeviceserviceSearchmodels(clt *core.SDKClient, req *tvupadmin.YunosOsupdateDeviceserviceSearchmodelsAPIRequest, resp *tvupadmin.YunosOsupdateDeviceserviceSearchmodelsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
