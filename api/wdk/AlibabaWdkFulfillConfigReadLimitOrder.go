package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfillconfigreadlimitorder 根据仓code查询仓限单配置
// alibaba.wdk.fulfill.config.read.limit.order
//
// 根据仓code查询仓限单配置
func Alibabawdkfulfillconfigreadlimitorder(clt *core.SDKClient, req *wdk.AlibabawdkfulfillconfigreadlimitorderAPIRequest, session string) (*wdk.AlibabawdkfulfillconfigreadlimitorderAPIResponse, error) {
	var resp wdk.AlibabawdkfulfillconfigreadlimitorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
