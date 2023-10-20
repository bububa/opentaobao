package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// Yunososupdatemodelsearch 机型检索
// yunos.osupdate.model.search
//
// 机型检索
func Yunososupdatemodelsearch(clt *core.SDKClient, req *tvupadmin.YunososupdatemodelsearchAPIRequest, session string) (*tvupadmin.YunososupdatemodelsearchAPIResponse, error) {
	var resp tvupadmin.YunososupdatemodelsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
