package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceupdateosstatus 更新系统版本审核状态
// yunos.tvpubadmin.device.updateosstatus
//
// 更新系统版本审核状态
func Yunostvpubadmindeviceupdateosstatus(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceupdateosstatusAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceupdateosstatusAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceupdateosstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
