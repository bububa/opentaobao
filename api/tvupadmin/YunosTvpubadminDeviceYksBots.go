package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvpubadminDeviceYksBots
获取设备列表
yunos.tvpubadmin.device.yks.bots

获取设备列表 */
func YunosTvpubadminDeviceYksBots(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksBotsAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceYksBotsAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminDeviceYksBotsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
