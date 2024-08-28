package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoJstAstrolabeStoreinventoryIteminitial 库存初始化接口
// taobao.jst.astrolabe.storeinventory.iteminitial
//
// ERP调用奇门的接口，对门店的库存进行初始化
func TaobaoJstAstrolabeStoreinventoryIteminitial(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest, resp *omniorder.TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
