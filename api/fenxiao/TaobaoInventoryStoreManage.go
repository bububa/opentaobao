package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoinventorystoremanage 创建或更新仓库
// taobao.inventory.store.manage
//
// 创建商家仓或者更新商家仓信息
func Taobaoinventorystoremanage(clt *core.SDKClient, req *fenxiao.TaobaoinventorystoremanageAPIRequest, session string) (*fenxiao.TaobaoinventorystoremanageAPIResponse, error) {
	var resp fenxiao.TaobaoinventorystoremanageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
