package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdateosfotaquery 系统升级分页查询
// yunos.osupdate.osfota.query
//
// 分页查询osoupdate系统升级列表
func Yunososupdateosfotaquery(clt *core.SDKClient, req *tvupadmin.YunososupdateosfotaqueryAPIRequest, session string) (*tvupadmin.YunososupdateosfotaqueryAPIResponse, error) {
	var resp tvupadmin.YunososupdateosfotaqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
