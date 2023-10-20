package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaposfundcashiershiftsummary 收银换班数据同步
// alibaba.pos.fund.cashier.shift.summary
//
// 收银换班数据同步，将每天收银换班的数据回流给商家。
func Alibabaposfundcashiershiftsummary(clt *core.SDKClient, req *wdk.AlibabaposfundcashiershiftsummaryAPIRequest, session string) (*wdk.AlibabaposfundcashiershiftsummaryAPIResponse, error) {
	var resp wdk.AlibabaposfundcashiershiftsummaryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
