package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenWarehouseinfoSynchronize 仓库同步接口
// taobao.qimen.warehouseinfo.synchronize
//
// 仓库同步接口
func TaobaoQimenWarehouseinfoSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenWarehouseinfoSynchronizeAPIRequest, resp *qimen.TaobaoQimenWarehouseinfoSynchronizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
