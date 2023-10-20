package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDiccontroltaskUpdate 停开服任务状态变更
// yunos.tvpubadmin.diccontroltask.update
//
// 停开服任务状态变更
func YunosTvpubadminDiccontroltaskUpdate(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDiccontroltaskUpdateAPIRequest, session string) (*tvupadmin.YunosTvpubadminDiccontroltaskUpdateAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminDiccontroltaskUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
