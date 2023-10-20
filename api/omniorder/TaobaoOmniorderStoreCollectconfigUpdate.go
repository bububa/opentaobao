package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstorecollectconfigupdate 门店自提配置修改
// taobao.omniorder.store.collectconfig.update
//
// 修改门店自提配置内容
func Taobaoomniorderstorecollectconfigupdate(clt *core.SDKClient, req *omniorder.TaobaoomniorderstorecollectconfigupdateAPIRequest, session string) (*omniorder.TaobaoomniorderstorecollectconfigupdateAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstorecollectconfigupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
