package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

/* TaobaoXhotelRatesUpdate
价格推送接口（批量全量）
taobao.xhotel.rates.update

酒店产品库rate批量更新房态信息 */
func TaobaoXhotelRatesUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRatesUpdateAPIRequest, session string) (*xhotelitem.TaobaoXhotelRatesUpdateAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelRatesUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
