package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmcustomeraddnew 导购域留资接口
// alibaba.lsy.crm.customer.add.new
//
// 导购域提供留资入口
func Alibabalsycrmcustomeraddnew(clt *core.SDKClient, req *tmallnr.AlibabalsycrmcustomeraddnewAPIRequest, session string) (*tmallnr.AlibabalsycrmcustomeraddnewAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmcustomeraddnewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
