package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindevicetvid 查询终端信息
// yunos.tvpubadmin.device.tvid
//
// 通过UUID查询终端信息
func Yunostvpubadmindevicetvid(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindevicetvidAPIRequest, session string) (*tvupadmin.YunostvpubadmindevicetvidAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindevicetvidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
