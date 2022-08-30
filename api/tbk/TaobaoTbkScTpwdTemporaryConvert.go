package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkScTpwdTemporaryConvert 淘宝客-服务商-淘口令解析&&转链（临时接口）
// taobao.tbk.sc.tpwd.temporary.convert
//
// 支持通过淘口令解析商品id，并提供对应的淘客转链接
func TaobaoTbkScTpwdTemporaryConvert(clt *core.SDKClient, req *tbk.TaobaoTbkScTpwdTemporaryConvertAPIRequest, session string) (*tbk.TaobaoTbkScTpwdTemporaryConvertAPIResponse, error) {
	var resp tbk.TaobaoTbkScTpwdTemporaryConvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
