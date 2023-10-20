package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdateappversionupdate 应用升级任务更新
// yunos.osupdate.appversion.update
//
// 应用升级任务更新
func Yunososupdateappversionupdate(clt *core.SDKClient, req *tvupadmin.YunososupdateappversionupdateAPIRequest, session string) (*tvupadmin.YunososupdateappversionupdateAPIResponse, error) {
	var resp tvupadmin.YunososupdateappversionupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
