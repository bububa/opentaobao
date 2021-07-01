package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

/* TaobaoXhotelRatesIncrement
价格推送接口（批量增量）
taobao.xhotel.rates.increment

Rate库存&价格增量更新接口，用户仅需要更新Rate中发生变化的库存日历&价格日历即可 */
func TaobaoXhotelRatesIncrement(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRatesIncrementAPIRequest, session string) (*xhotelitem.TaobaoXhotelRatesIncrementAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelRatesIncrementAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
