package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunostvpubadminuserrights 获取用户权益
// yunos.tvpubadmin.user.rights
//
// 获取用户权益
func Yunostvpubadminuserrights(clt *core.SDKClient, req *tvupadmin.YunostvpubadminuserrightsAPIRequest, session string) (*tvupadmin.YunostvpubadminuserrightsAPIResponse, error) {
	var resp tvupadmin.YunostvpubadminuserrightsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
