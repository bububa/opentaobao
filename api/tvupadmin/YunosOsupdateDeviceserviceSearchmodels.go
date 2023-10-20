package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdatedeviceservicesearchmodels 根据关键词检索设备型号
// yunos.osupdate.deviceservice.searchmodels
//
// 根据关键词检索设备型号
func Yunososupdatedeviceservicesearchmodels(clt *core.SDKClient, req *tvupadmin.YunososupdatedeviceservicesearchmodelsAPIRequest, session string) (*tvupadmin.YunososupdatedeviceservicesearchmodelsAPIResponse, error) {
	var resp tvupadmin.YunososupdatedeviceservicesearchmodelsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
