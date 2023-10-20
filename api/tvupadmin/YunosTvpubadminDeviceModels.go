package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceModels 获取品牌下设备列表
// yunos.tvpubadmin.device.models
//
// 获取品牌下设备列表
func YunosTvpubadminDeviceModels(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceModelsAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceModelsAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminDeviceModelsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
