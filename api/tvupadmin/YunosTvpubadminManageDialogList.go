package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminManageDialogList 分页获取弹窗列表
// yunos.tvpubadmin.manage.dialog.list
//
// 分页获取弹窗配置列表
func YunosTvpubadminManageDialogList(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogListAPIRequest, resp *tvupadmin.YunosTvpubadminManageDialogListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
