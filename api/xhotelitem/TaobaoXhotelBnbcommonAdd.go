package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbcommonAdd 通用调用top接口
// taobao.xhotel.bnbcommon.add
//
// 通用调用top接口
func TaobaoXhotelBnbcommonAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbcommonAddAPIRequest, resp *xhotelitem.TaobaoXhotelBnbcommonAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
