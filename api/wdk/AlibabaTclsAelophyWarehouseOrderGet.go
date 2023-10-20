package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophywarehouseorderget 仓作业单获取
// alibaba.tcls.aelophy.warehouse.order.get
//
// 仓作业单获取
func Alibabatclsaelophywarehouseorderget(clt *core.SDKClient, req *wdk.AlibabatclsaelophywarehouseordergetAPIRequest, session string) (*wdk.AlibabatclsaelophywarehouseordergetAPIResponse, error) {
	var resp wdk.AlibabatclsaelophywarehouseordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
