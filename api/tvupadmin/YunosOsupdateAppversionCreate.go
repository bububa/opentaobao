package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdateappversioncreate 创建应用升级任务
// yunos.osupdate.appversion.create
//
// 创建应用升级任务
func Yunososupdateappversioncreate(clt *core.SDKClient, req *tvupadmin.YunososupdateappversioncreateAPIRequest, session string) (*tvupadmin.YunososupdateappversioncreateAPIResponse, error) {
	var resp tvupadmin.YunososupdateappversioncreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
