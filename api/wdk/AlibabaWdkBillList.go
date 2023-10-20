package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkbilllist 五道口账单拉取接口
// alibaba.wdk.bill.list
//
// 五道口账单拉取接口
func Alibabawdkbilllist(clt *core.SDKClient, req *wdk.AlibabawdkbilllistAPIRequest, session string) (*wdk.AlibabawdkbilllistAPIResponse, error) {
	var resp wdk.AlibabawdkbilllistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
