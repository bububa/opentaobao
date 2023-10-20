package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenwarehouseinfosynchronize 仓库同步接口
// taobao.qimen.warehouseinfo.synchronize
//
// 仓库同步接口
func Taobaoqimenwarehouseinfosynchronize(clt *core.SDKClient, req *qimen.TaobaoqimenwarehouseinfosynchronizeAPIRequest, session string) (*qimen.TaobaoqimenwarehouseinfosynchronizeAPIResponse, error) {
	var resp qimen.TaobaoqimenwarehouseinfosynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
