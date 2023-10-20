package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceModels 获取品牌下设备列表
// yunos.tvpubadmin.device.models
//
// 获取品牌下设备列表
func YunosTvpubadminDeviceModels(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceModelsAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceModelsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
