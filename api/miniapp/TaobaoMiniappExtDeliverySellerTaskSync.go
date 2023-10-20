package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaominiappextdeliverysellertasksync 同步外投任务
// taobao.miniapp.ext.delivery.seller.task.sync
//
// 同步外投任务
func Taobaominiappextdeliverysellertasksync(clt *core.SDKClient, req *miniapp.TaobaominiappextdeliverysellertasksyncAPIRequest, session string) (*miniapp.TaobaominiappextdeliverysellertasksyncAPIResponse, error) {
	var resp miniapp.TaobaominiappextdeliverysellertasksyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
