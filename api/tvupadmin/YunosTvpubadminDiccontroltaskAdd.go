package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDiccontroltaskAdd 新增停开服任务
// yunos.tvpubadmin.diccontroltask.add
//
// 新增停开服任务
func YunosTvpubadminDiccontroltaskAdd(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDiccontroltaskAddAPIRequest, session string) (*tvupadmin.YunosTvpubadminDiccontroltaskAddAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminDiccontroltaskAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
