package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// Taobaominiappadvancedtradeinfopricemodify 高级定制商家传入改价信息
// taobao.miniapp.advanced.tradeinfo.price.modify
//
// 高级定制商家传入改价信息
func Taobaominiappadvancedtradeinfopricemodify(clt *core.SDKClient, req *opentrade.TaobaominiappadvancedtradeinfopricemodifyAPIRequest, session string) (*opentrade.TaobaominiappadvancedtradeinfopricemodifyAPIResponse, error) {
	var resp opentrade.TaobaominiappadvancedtradeinfopricemodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
