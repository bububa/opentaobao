package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagedialogadd 新增全局弹窗
// yunos.tvpubadmin.manage.dialog.add
//
// 新增全局弹窗
func Yunostvpubadminmanagedialogadd(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagedialogaddAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagedialogaddAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagedialogaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
