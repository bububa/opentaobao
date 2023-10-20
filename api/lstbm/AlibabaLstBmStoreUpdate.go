package lstbm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstbm"
)

// Alibabalstbmstoreupdate 修改品牌商自有门店数据
// alibaba.lst.bm.store.update
//
// 修改品牌商自有门店数据
func Alibabalstbmstoreupdate(clt *core.SDKClient, req *lstbm.AlibabalstbmstoreupdateAPIRequest, session string) (*lstbm.AlibabalstbmstoreupdateAPIResponse, error) {
	var resp lstbm.AlibabalstbmstoreupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
