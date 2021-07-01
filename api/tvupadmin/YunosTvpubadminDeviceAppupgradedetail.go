package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvpubadminDeviceAppupgradedetail
获取应用升级详情
yunos.tvpubadmin.device.appupgradedetail

获取应用升级详情 */
func YunosTvpubadminDeviceAppupgradedetail(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceAppupgradedetailAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceAppupgradedetailAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminDeviceAppupgradedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
