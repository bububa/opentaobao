package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// TaobaoXhotelComboStatusGet 酒店宝贝状态查询
// taobao.xhotel.combo.status.get
//
// 酒店宝贝状态查询
func TaobaoXhotelComboStatusGet(clt *core.SDKClient, req *tuanhotel.TaobaoXhotelComboStatusGetAPIRequest, resp *tuanhotel.TaobaoXhotelComboStatusGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
