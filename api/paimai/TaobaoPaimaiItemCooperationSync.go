package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoPaimaiItemCooperationSync 商品同步
// taobao.paimai.item.cooperation.sync
//
// 商品同步
func TaobaoPaimaiItemCooperationSync(clt *core.SDKClient, req *paimai.TaobaoPaimaiItemCooperationSyncAPIRequest, session string) (*paimai.TaobaoPaimaiItemCooperationSyncAPIResponse, error) {
	var resp paimai.TaobaoPaimaiItemCooperationSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
