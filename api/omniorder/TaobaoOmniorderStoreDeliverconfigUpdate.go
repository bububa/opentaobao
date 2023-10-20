package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstoredeliverconfigupdate 修改门店发货配置内容
// taobao.omniorder.store.deliverconfig.update
//
// 修改门店发货配置内容
func Taobaoomniorderstoredeliverconfigupdate(clt *core.SDKClient, req *omniorder.TaobaoomniorderstoredeliverconfigupdateAPIRequest, session string) (*omniorder.TaobaoomniorderstoredeliverconfigupdateAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstoredeliverconfigupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
