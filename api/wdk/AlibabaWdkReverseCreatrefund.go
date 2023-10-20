package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkreversecreatrefund 逆向提交
// alibaba.wdk.reverse.creatrefund
//
// 逆向申请提交
func Alibabawdkreversecreatrefund(clt *core.SDKClient, req *wdk.AlibabawdkreversecreatrefundAPIRequest, session string) (*wdk.AlibabawdkreversecreatrefundAPIResponse, error) {
	var resp wdk.AlibabawdkreversecreatrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
