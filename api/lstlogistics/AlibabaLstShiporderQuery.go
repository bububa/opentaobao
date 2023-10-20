package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// Alibabalstshiporderquery 零售通发货单查询
// alibaba.lst.shiporder.query
//
// 通过该接口可以查询零售通运保保发货单，并处理相关业务流程。
func Alibabalstshiporderquery(clt *core.SDKClient, req *lstlogistics.AlibabalstshiporderqueryAPIRequest, session string) (*lstlogistics.AlibabalstshiporderqueryAPIResponse, error) {
	var resp lstlogistics.AlibabalstshiporderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
