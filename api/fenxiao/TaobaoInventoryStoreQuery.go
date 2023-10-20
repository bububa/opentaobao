package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoinventorystorequery 查询仓库信息
// taobao.inventory.store.query
//
// 查询商家仓信息
func Taobaoinventorystorequery(clt *core.SDKClient, req *fenxiao.TaobaoinventorystorequeryAPIRequest, session string) (*fenxiao.TaobaoinventorystorequeryAPIResponse, error) {
	var resp fenxiao.TaobaoinventorystorequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
