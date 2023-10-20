package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceyksskillmodify 修改技能
// yunos.tvpubadmin.device.yks.skill.modify
//
// 修改技能
func Yunostvpubadmindeviceyksskillmodify(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceyksskillmodifyAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceyksskillmodifyAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceyksskillmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
