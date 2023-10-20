package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdktimeget 获得当前系统时间
// alibaba.wdk.time.get
//
// 获得当前系统时间
func Alibabawdktimeget(clt *core.SDKClient, req *wdk.AlibabawdktimegetAPIRequest, session string) (*wdk.AlibabawdktimegetAPIResponse, error) {
	var resp wdk.AlibabawdktimegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
