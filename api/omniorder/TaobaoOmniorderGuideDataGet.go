package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderguidedataget 获取全渠道导购产品数据
// taobao.omniorder.guide.data.get
//
// 获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
func Taobaoomniorderguidedataget(clt *core.SDKClient, req *omniorder.TaobaoomniorderguidedatagetAPIRequest, session string) (*omniorder.TaobaoomniorderguidedatagetAPIResponse, error) {
	var resp omniorder.TaobaoomniorderguidedatagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
