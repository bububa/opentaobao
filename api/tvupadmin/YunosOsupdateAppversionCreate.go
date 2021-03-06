package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateAppversionCreate 创建应用升级任务
// yunos.osupdate.appversion.create
//
// 创建应用升级任务
func YunosOsupdateAppversionCreate(clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionCreateAPIRequest, session string) (*tvupadmin.YunosOsupdateAppversionCreateAPIResponse, error) {
	var resp tvupadmin.YunosOsupdateAppversionCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
