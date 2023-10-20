package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdateosfotapublish 系统升级发布
// yunos.osupdate.osfota.publish
//
// 发布osupdate系统升级任务
func Yunososupdateosfotapublish(clt *core.SDKClient, req *tvupadmin.YunososupdateosfotapublishAPIRequest, session string) (*tvupadmin.YunososupdateosfotapublishAPIResponse, error) {
	var resp tvupadmin.YunososupdateosfotapublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
