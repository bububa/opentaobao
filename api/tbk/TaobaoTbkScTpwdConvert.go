package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScTpwdConvert 淘宝客-服务商-淘口令解析&转链
// taobao.tbk.sc.tpwd.convert
//
// 支持通过淘口令解析商品id，并提供对应的淘客转链接
func TaobaoTbkScTpwdConvert(clt *core.SDKClient, req *tbk.TaobaoTbkScTpwdConvertAPIRequest, session string) (*tbk.TaobaoTbkScTpwdConvertAPIResponse, error) {
	var resp tbk.TaobaoTbkScTpwdConvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
