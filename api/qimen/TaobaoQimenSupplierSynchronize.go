package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenSupplierSynchronize 供应商同步接口
// taobao.qimen.supplier.synchronize
//
// 这个接口用来同步供应商信息
func TaobaoQimenSupplierSynchronize(clt *core.SDKClient, req *qimen.TaobaoQimenSupplierSynchronizeAPIRequest, resp *qimen.TaobaoQimenSupplierSynchronizeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
