package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbImportsResourceTransferstoreGet 根据指定的资源获取所有中转仓列表
// taobao.wlb.imports.resource.transferstore.get
//
// 根据指定的资源获取所有中转仓列表
func TaobaoWlbImportsResourceTransferstoreGet(clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsResourceTransferstoreGetAPIRequest, session string) (*wlbimports.TaobaoWlbImportsResourceTransferstoreGetAPIResponse, error) {
	var resp wlbimports.TaobaoWlbImportsResourceTransferstoreGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
