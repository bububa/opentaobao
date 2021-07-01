package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvpubadminManageDialogList
分页获取弹窗列表
yunos.tvpubadmin.manage.dialog.list

分页获取弹窗配置列表 */
func YunosTvpubadminManageDialogList(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminManageDialogListAPIRequest, session string) (*tvupadmin.YunosTvpubadminManageDialogListAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminManageDialogListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
