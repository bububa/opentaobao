package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosOsupdateModelSearch 机型检索
// yunos.osupdate.model.search
//
// 机型检索
func YunosOsupdateModelSearch(clt *core.SDKClient, req *tvupadmin.YunosOsupdateModelSearchAPIRequest, session string) (*tvupadmin.YunosOsupdateModelSearchAPIResponse, error) {
	var resp tvupadmin.YunosOsupdateModelSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
