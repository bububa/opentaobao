package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdateosfotaadd 添加系统升级任务
// yunos.osupdate.osfota.add
//
// 添加osupdate系统升级任务
func Yunososupdateosfotaadd(clt *core.SDKClient, req *tvupadmin.YunososupdateosfotaaddAPIRequest, session string) (*tvupadmin.YunososupdateosfotaaddAPIResponse, error) {
	var resp tvupadmin.YunososupdateosfotaaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
