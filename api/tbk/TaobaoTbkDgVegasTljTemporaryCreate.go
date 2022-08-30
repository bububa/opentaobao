package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgVegasTljTemporaryCreate 淘宝客-推广者-淘礼金创建测试营销ID
// taobao.tbk.dg.vegas.tlj.temporary.create
//
// 淘宝客-推广者-淘礼金创建测试营销ID
func TaobaoTbkDgVegasTljTemporaryCreate(clt *core.SDKClient, req *tbk.TaobaoTbkDgVegasTljTemporaryCreateAPIRequest, session string) (*tbk.TaobaoTbkDgVegasTljTemporaryCreateAPIResponse, error) {
	var resp tbk.TaobaoTbkDgVegasTljTemporaryCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
