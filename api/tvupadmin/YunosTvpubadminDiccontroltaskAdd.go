package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindiccontroltaskadd 新增停开服任务
// yunos.tvpubadmin.diccontroltask.add
//
// 新增停开服任务
func Yunostvpubadmindiccontroltaskadd(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindiccontroltaskaddAPIRequest, session string) (*tvupadmin.YunostvpubadmindiccontroltaskaddAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindiccontroltaskaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
