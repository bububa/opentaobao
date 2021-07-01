package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

/* TaobaoVmarketEticketPackageBaseGet
获取包基本信息
taobao.vmarket.eticket.package.base.get

获取包基本信息 */
func TaobaoVmarketEticketPackageBaseGet(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketPackageBaseGetAPIRequest, session string) (*eticket.TaobaoVmarketEticketPackageBaseGetAPIResponse, error) {
	var resp eticket.TaobaoVmarketEticketPackageBaseGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
