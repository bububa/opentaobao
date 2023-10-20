package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbksctpwdconvert 淘宝客-服务商-淘口令解析&转链
// taobao.tbk.sc.tpwd.convert
//
// 支持通过淘口令解析商品id，并提供对应的淘客转链接
func Taobaotbksctpwdconvert(clt *core.SDKClient, req *tbk.TaobaotbksctpwdconvertAPIRequest, session string) (*tbk.TaobaotbksctpwdconvertAPIResponse, error) {
	var resp tbk.TaobaotbksctpwdconvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
