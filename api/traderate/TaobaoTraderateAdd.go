package traderate

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traderate"
)

/* TaobaoTraderateAdd
新增单个评价
taobao.traderate.add

新增单个评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不能再通过该接口进行评价</font>) */
func TaobaoTraderateAdd(clt *core.SDKClient, req *traderate.TaobaoTraderateAddAPIRequest, session string) (*traderate.TaobaoTraderateAddAPIResponse, error) {
	var resp traderate.TaobaoTraderateAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
