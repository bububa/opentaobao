package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdateappversionpublish 发布应用升级
// yunos.osupdate.appversion.publish
//
// 发布应用升级任务
func Yunososupdateappversionpublish(clt *core.SDKClient, req *tvupadmin.YunososupdateappversionpublishAPIRequest, session string) (*tvupadmin.YunososupdateappversionpublishAPIResponse, error) {
	var resp tvupadmin.YunososupdateappversionpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
