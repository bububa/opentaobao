package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltradeaddressparse openmall服务地址区域码解析
// taobao.openmall.trade.address.parse
//
// openmall服务，解析地址区域码，获取创建订单等接口中的区域码信息
func Taobaoopenmalltradeaddressparse(clt *core.SDKClient, req *openmall.TaobaoopenmalltradeaddressparseAPIRequest, session string) (*openmall.TaobaoopenmalltradeaddressparseAPIResponse, error) {
	var resp openmall.TaobaoopenmalltradeaddressparseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
