package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryStoreManage 创建或更新仓库
// taobao.inventory.store.manage
//
// 创建商家仓或者更新商家仓信息
func TaobaoInventoryStoreManage(clt *core.SDKClient, req *fenxiao.TaobaoInventoryStoreManageAPIRequest, session string) (*fenxiao.TaobaoInventoryStoreManageAPIResponse, error) {
	var resp fenxiao.TaobaoInventoryStoreManageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
