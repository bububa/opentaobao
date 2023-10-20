package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoregionpricemanage 编辑区域价格
// taobao.region.price.manage
//
// 编辑区域价格
func Taobaoregionpricemanage(clt *core.SDKClient, req *fenxiao.TaobaoregionpricemanageAPIRequest, session string) (*fenxiao.TaobaoregionpricemanageAPIResponse, error) {
	var resp fenxiao.TaobaoregionpricemanageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
