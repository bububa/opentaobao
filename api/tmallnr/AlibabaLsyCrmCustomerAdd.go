package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Alibabalsycrmcustomeradd 私域导购添加活动留资入口
// alibaba.lsy.crm.customer.add
//
// 私域导购添加活动留资入口
func Alibabalsycrmcustomeradd(clt *core.SDKClient, req *tmallnr.AlibabalsycrmcustomeraddAPIRequest, session string) (*tmallnr.AlibabalsycrmcustomeraddAPIResponse, error) {
	var resp tmallnr.AlibabalsycrmcustomeraddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
