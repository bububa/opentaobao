package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceUpdateosstatus 更新系统版本审核状态
// yunos.tvpubadmin.device.updateosstatus
//
// 更新系统版本审核状态
func YunosTvpubadminDeviceUpdateosstatus(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceUpdateosstatusAPIRequest, session string) (*tvupadmin.YunosTvpubadminDeviceUpdateosstatusAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminDeviceUpdateosstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
