package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminUserRights 获取用户权益
// yunos.tvpubadmin.user.rights
//
// 获取用户权益
func YunosTvpubadminUserRights(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminUserRightsAPIRequest, session string) (*tvupadmin.YunosTvpubadminUserRightsAPIResponse, error) {
	var resp tvupadmin.YunosTvpubadminUserRightsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
