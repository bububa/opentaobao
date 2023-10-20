package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoregionpricequery 区域价格查询
// taobao.region.price.query
//
// 区域价格查询
func Taobaoregionpricequery(clt *core.SDKClient, req *fenxiao.TaobaoregionpricequeryAPIRequest, session string) (*fenxiao.TaobaoregionpricequeryAPIResponse, error) {
	var resp fenxiao.TaobaoregionpricequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
