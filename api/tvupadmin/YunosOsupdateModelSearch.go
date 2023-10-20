package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateModelSearch 机型检索
// yunos.osupdate.model.search
//
// 机型检索
func YunosOsupdateModelSearch(clt *core.SDKClient, req *tvupadmin.YunosOsupdateModelSearchAPIRequest, resp *tvupadmin.YunosOsupdateModelSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
