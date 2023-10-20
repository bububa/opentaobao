package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdateappversionquery 分页获取桌面升级任务
// yunos.osupdate.appversion.query
//
// 分页获取桌面升级任务
func Yunososupdateappversionquery(clt *core.SDKClient, req *tvupadmin.YunososupdateappversionqueryAPIRequest, session string) (*tvupadmin.YunososupdateappversionqueryAPIResponse, error) {
	var resp tvupadmin.YunososupdateappversionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
