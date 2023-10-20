package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkelemebillget 饿了么日维度对账单查询
// alibaba.wdk.eleme.bill.get
//
// 查询饿了么日维度对账单信息
func Alibabawdkelemebillget(clt *core.SDKClient, req *wdk.AlibabawdkelemebillgetAPIRequest, session string) (*wdk.AlibabawdkelemebillgetAPIResponse, error) {
	var resp wdk.AlibabawdkelemebillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
