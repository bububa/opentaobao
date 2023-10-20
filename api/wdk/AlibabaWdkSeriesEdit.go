package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkseriesedit 系列品变更-更新系列
// alibaba.wdk.series.edit
//
// 系列品变更-更新系列
func Alibabawdkseriesedit(clt *core.SDKClient, req *wdk.AlibabawdkserieseditAPIRequest, session string) (*wdk.AlibabawdkserieseditAPIResponse, error) {
	var resp wdk.AlibabawdkserieseditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
