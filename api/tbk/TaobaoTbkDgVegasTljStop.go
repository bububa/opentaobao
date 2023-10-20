package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgVegasTljStop 淘宝客-推广者-淘礼金暂停发放
// taobao.tbk.dg.vegas.tlj.stop
//
// 淘宝客推广者可对已经创建的淘礼金暂停发放
func TaobaoTbkDgVegasTljStop(clt *core.SDKClient, req *tbk.TaobaoTbkDgVegasTljStopAPIRequest, session string) (*tbk.TaobaoTbkDgVegasTljStopAPIResponse, error) {
	var resp tbk.TaobaoTbkDgVegasTljStopAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
