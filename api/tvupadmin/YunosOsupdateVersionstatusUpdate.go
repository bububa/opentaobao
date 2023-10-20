package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdateversionstatusupdate 更新应用升级状态
// yunos.osupdate.versionstatus.update
//
// 更新应用升级状态
func Yunososupdateversionstatusupdate(clt *core.SDKClient, req *tvupadmin.YunososupdateversionstatusupdateAPIRequest, session string) (*tvupadmin.YunososupdateversionstatusupdateAPIResponse, error) {
	var resp tvupadmin.YunososupdateversionstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
