package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Taobaophonebankcreditprocess 虚拟话费任务银行信用卡办理进度回传
// taobao.phone.bank.credit.process
//
// 虚拟话费任务银行信用卡办理进度回传
func Taobaophonebankcreditprocess(clt *core.SDKClient, req *alicom.TaobaophonebankcreditprocessAPIRequest, session string) (*alicom.TaobaophonebankcreditprocessAPIResponse, error) {
	var resp alicom.TaobaophonebankcreditprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
