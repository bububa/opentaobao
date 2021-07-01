package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvpubadminContentDeviceGetvendor
查询设备Vendor信息
yunos.tvpubadmin.content.device.getvendor

查询设备Vendor信息 */
func YunosTvpubadminContentDeviceGetvendor(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentDeviceGetvendorAPIRequest, session string) (*tvupadmin.YunosTvpubadminContentDeviceGetvendorAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminContentDeviceGetvendorAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
