package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

/* TaobaoOcOrderApUpdate
按OC订单分账
taobao.oc.order.ap.update

对OC订单执行分账操作 */
func TaobaoOcOrderApUpdate(clt *core.SDKClient, req *jst.TaobaoOcOrderApUpdateAPIRequest, session string) (*jst.TaobaoOcOrderApUpdateAPIResponse, error) {
	var resp jst.TaobaoOcOrderApUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
