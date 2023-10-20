package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceupdateappstatus 更新应用版本审核状态
// yunos.tvpubadmin.device.updateappstatus
//
// 更新应用版本审核状态
func Yunostvpubadmindeviceupdateappstatus(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceupdateappstatusAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceupdateappstatusAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceupdateappstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
