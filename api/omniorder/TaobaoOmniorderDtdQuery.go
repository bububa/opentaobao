package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderdtdquery 门店自送根据核销码查订单
// taobao.omniorder.dtd.query
//
// 门店自送根据核销码码查询订单信息
func Taobaoomniorderdtdquery(clt *core.SDKClient, req *omniorder.TaobaoomniorderdtdqueryAPIRequest, session string) (*omniorder.TaobaoomniorderdtdqueryAPIResponse, error) {
	var resp omniorder.TaobaoomniorderdtdqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
