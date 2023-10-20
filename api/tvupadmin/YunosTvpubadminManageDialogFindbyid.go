package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagedialogfindbyid 根据id查询全局弹窗
// yunos.tvpubadmin.manage.dialog.findbyid
//
// 根据id查询全局弹窗
func Yunostvpubadminmanagedialogfindbyid(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagedialogfindbyidAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagedialogfindbyidAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagedialogfindbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
