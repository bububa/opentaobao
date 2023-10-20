package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagedialogdelete 删除全局弹窗
// yunos.tvpubadmin.manage.dialog.delete
//
// 删除全局弹窗
func Yunostvpubadminmanagedialogdelete(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagedialogdeleteAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagedialogdeleteAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagedialogdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
