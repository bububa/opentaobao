package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindiccontroltaskupdate 停开服任务状态变更
// yunos.tvpubadmin.diccontroltask.update
//
// 停开服任务状态变更
func Yunostvpubadmindiccontroltaskupdate(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindiccontroltaskupdateAPIRequest, session string) (*tvupadmin.YunostvpubadmindiccontroltaskupdateAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindiccontroltaskupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
