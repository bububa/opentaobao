package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// Alibabalstlogisticsthirdpartcompanylist 供应商-异云-第三方物流公司列表
// alibaba.lst.logistics.thirdpart.company.list
//
// 异地云仓发货时，需填写的第三方物流公司列表
func Alibabalstlogisticsthirdpartcompanylist(clt *core.SDKClient, req *lstlogistics.AlibabalstlogisticsthirdpartcompanylistAPIRequest, session string) (*lstlogistics.AlibabalstlogisticsthirdpartcompanylistAPIResponse, error) {
	var resp lstlogistics.AlibabalstlogisticsthirdpartcompanylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
