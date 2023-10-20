package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindeviceyksskilldelete 技能删除
// yunos.tvpubadmin.device.yks.skill.delete
//
// 删除技能
func Yunostvpubadmindeviceyksskilldelete(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindeviceyksskilldeleteAPIRequest, session string) (*tvupadmin.YunostvpubadmindeviceyksskilldeleteAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindeviceyksskilldeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
