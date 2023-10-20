package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderstoresdtquerystation 速店通查询站点信息
// taobao.omniorder.store.sdtquerystation
//
// 速店通查询站点信息
func Taobaoomniorderstoresdtquerystation(clt *core.SDKClient, req *omniorder.TaobaoomniorderstoresdtquerystationAPIRequest, session string) (*omniorder.TaobaoomniorderstoresdtquerystationAPIResponse, error) {
	var resp omniorder.TaobaoomniorderstoresdtquerystationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
