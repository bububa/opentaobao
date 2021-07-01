package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvpubadminDeviceApks
获取停开服apk列表
yunos.tvpubadmin.device.apks

获取停开服apk列表 */
func YunosTvpubadminDeviceApks(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceApksAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceApksAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminDeviceApksAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
