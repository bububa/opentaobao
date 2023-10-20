package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimensuppliersynchronize 供应商同步接口
// taobao.qimen.supplier.synchronize
//
// 这个接口用来同步供应商信息
func Taobaoqimensuppliersynchronize(clt *core.SDKClient, req *qimen.TaobaoqimensuppliersynchronizeAPIRequest, session string) (*qimen.TaobaoqimensuppliersynchronizeAPIResponse, error) {
	var resp qimen.TaobaoqimensuppliersynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
