package lstbm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstbm"
)

// Alibabalstbmstoreadd 导入品牌商自有门店
// alibaba.lst.bm.store.add
//
// 导入品牌商自有门店
func Alibabalstbmstoreadd(clt *core.SDKClient, req *lstbm.AlibabalstbmstoreaddAPIRequest, session string) (*lstbm.AlibabalstbmstoreaddAPIResponse, error) {
	var resp lstbm.AlibabalstbmstoreaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
