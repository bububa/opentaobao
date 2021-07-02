package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateDeviceserviceSearchmodels 根据关键词检索设备型号
// yunos.osupdate.deviceservice.searchmodels
//
// 根据关键词检索设备型号
func YunosOsupdateDeviceserviceSearchmodels(clt *core.SDKClient, req *tvupadmin.YunosOsupdateDeviceserviceSearchmodelsAPIRequest, session string) (*tvupadmin.YunosOsupdateDeviceserviceSearchmodelsAPIResponse, error) {
	var resp tvupadmin.YunosOsupdateDeviceserviceSearchmodelsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
