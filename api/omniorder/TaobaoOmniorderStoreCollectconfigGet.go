package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstorecollectconfigget 查询门店自提配置内容
// taobao.omniorder.store.collectconfig.get
//
// 查询门店自提配置内容
func Taobaoomniorderstorecollectconfigget(clt *core.SDKClient, req *omniorder.TaobaoomniorderstorecollectconfiggetAPIRequest, session string) (*omniorder.TaobaoomniorderstorecollectconfiggetAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstorecollectconfiggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
