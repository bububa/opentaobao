package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Taobaophonebankcreditcheck 虚拟话费任务银行信用卡办理检查
// taobao.phone.bank.credit.check
//
// 虚拟话费任务银行信用卡办理检查
func Taobaophonebankcreditcheck(clt *core.SDKClient, req *alicom.TaobaophonebankcreditcheckAPIRequest, session string) (*alicom.TaobaophonebankcreditcheckAPIResponse, error) {
	var resp alicom.TaobaophonebankcreditcheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
