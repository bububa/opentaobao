package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadmindiccontroltaskgetinfo 获取停开服任务详情
// yunos.tvpubadmin.diccontroltask.getinfo
//
// 获取停开服任务详情
func Yunostvpubadmindiccontroltaskgetinfo(clt *core.SDKClient, req *tvupadmin.YunostvpubadmindiccontroltaskgetinfoAPIRequest, session string) (*tvupadmin.YunostvpubadmindiccontroltaskgetinfoAPIResponse, error) {
	var resp tvupadmin.YunostvpubadmindiccontroltaskgetinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
