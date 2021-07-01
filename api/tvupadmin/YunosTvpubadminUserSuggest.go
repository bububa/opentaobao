package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

/* YunosTvpubadminUserSuggest
获取关联账户列表
yunos.tvpubadmin.user.suggest

获取关联账户列表 */
func YunosTvpubadminUserSuggest(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminUserSuggestAPIRequest, session string) (*tvupadmin.YunosTvpubadminUserSuggestAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminUserSuggestAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
