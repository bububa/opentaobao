package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkSpreadGet 淘宝客-公用-长链转短链
// taobao.tbk.spread.get
//
// 输入一个原始的链接，转换得到指定的传播方式，如二维码，淘口令，短连接；
// 现阶段只支持短连接。
func TaobaoTbkSpreadGet(clt *core.SDKClient, req *tbk.TaobaoTbkSpreadGetAPIRequest, session string) (*tbk.TaobaoTbkSpreadGetAPIResponse, error) {
	var resp tbk.TaobaoTbkSpreadGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
