package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkumsshiftget 移库单获取
// alibaba.wdk.ums.shift.get
//
// 移库单获取
func Alibabawdkumsshiftget(clt *core.SDKClient, req *wdk.AlibabawdkumsshiftgetAPIRequest, session string) (*wdk.AlibabawdkumsshiftgetAPIResponse, error) {
	var resp wdk.AlibabawdkumsshiftgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
