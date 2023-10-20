package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagedialogedit 编辑全局弹窗
// yunos.tvpubadmin.manage.dialog.edit
//
// 编辑全局弹窗
func Yunostvpubadminmanagedialogedit(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagedialogeditAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagedialogeditAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagedialogeditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
