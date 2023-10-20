package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// Alibabalstshipordercancel 零售通发货单取消
// alibaba.lst.shiporder.cancel
//
// 通过该接口可以取消零售通运保保发货单，并处理相关业务流程。
func Alibabalstshipordercancel(clt *core.SDKClient, req *lstlogistics.AlibabalstshipordercancelAPIRequest, session string) (*lstlogistics.AlibabalstshipordercancelAPIResponse, error) {
	var resp lstlogistics.AlibabalstshipordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
