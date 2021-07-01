package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

/* TaobaoXhotelRateUpdate
价格推送接口（全量更新）
taobao.xhotel.rate.update

酒店产品库rate更新 */
func TaobaoXhotelRateUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateUpdateAPIRequest, session string) (*xhotelitem.TaobaoXhotelRateUpdateAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelRateUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
