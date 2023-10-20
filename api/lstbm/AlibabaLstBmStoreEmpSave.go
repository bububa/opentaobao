package lstbm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstbm"
)

// Alibabalstbmstoreempsave 保存品牌商自有门店和内部业代之间的关系
// alibaba.lst.bm.store.emp.save
//
// 保存品牌商自有门店和内部业代之间的关系
func Alibabalstbmstoreempsave(clt *core.SDKClient, req *lstbm.AlibabalstbmstoreempsaveAPIRequest, session string) (*lstbm.AlibabalstbmstoreempsaveAPIResponse, error) {
	var resp lstbm.AlibabalstbmstoreempsaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
