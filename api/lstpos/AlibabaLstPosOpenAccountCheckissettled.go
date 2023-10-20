package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// Alibabalstposopenaccountcheckissettled 校验当前用户是否入驻了零售通门店接口
// alibaba.lst.pos.open.account.checkissettled
//
// 校验当前用户是否入驻了零售通门店接口
func Alibabalstposopenaccountcheckissettled(clt *core.SDKClient, req *lstpos.AlibabalstposopenaccountcheckissettledAPIRequest, session string) (*lstpos.AlibabalstposopenaccountcheckissettledAPIResponse, error) {
	var resp lstpos.AlibabalstposopenaccountcheckissettledAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
