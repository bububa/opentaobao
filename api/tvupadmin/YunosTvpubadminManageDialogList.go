package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminmanagedialoglist 分页获取弹窗列表
// yunos.tvpubadmin.manage.dialog.list
//
// 分页获取弹窗配置列表
func Yunostvpubadminmanagedialoglist(clt *core.SDKClient, req *tvupadmin.YunostvpubadminmanagedialoglistAPIRequest, session string) (*tvupadmin.YunostvpubadminmanagedialoglistAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminmanagedialoglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
