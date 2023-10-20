package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkreversecreatefeatch 逆向取货
// alibaba.wdk.reverse.createfeatch
//
// 发起逆向取货
func Alibabawdkreversecreatefeatch(clt *core.SDKClient, req *wdk.AlibabawdkreversecreatefeatchAPIRequest, session string) (*wdk.AlibabawdkreversecreatefeatchAPIResponse, error) {
	var resp wdk.AlibabawdkreversecreatefeatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
