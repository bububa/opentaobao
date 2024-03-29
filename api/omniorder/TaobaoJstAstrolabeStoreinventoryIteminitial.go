package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoJstAstrolabeStoreinventoryIteminitial 库存初始化接口
// taobao.jst.astrolabe.storeinventory.iteminitial
//
// ERP调用奇门的接口，对门店的库存进行初始化
func TaobaoJstAstrolabeStoreinventoryIteminitial(clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryIteminitialAPIRequest, resp *omniorder.TaobaoJstAstrolabeStoreinventoryIteminitialAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
