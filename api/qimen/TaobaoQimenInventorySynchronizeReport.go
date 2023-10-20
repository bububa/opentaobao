package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimeninventorysynchronizereport 库存状态同步确认接口
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
func Taobaoqimeninventorysynchronizereport(clt *core.SDKClient, req *qimen.TaobaoqimeninventorysynchronizereportAPIRequest, session string) (*qimen.TaobaoqimeninventorysynchronizereportAPIResponse, error) {
	var resp qimen.TaobaoqimeninventorysynchronizereportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
