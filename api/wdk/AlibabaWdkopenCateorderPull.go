package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkopencateorderpull 商户回传餐饮加工单状态
// alibaba.wdkopen.cateorder.pull
//
// 商户回传餐饮加工单状态
func Alibabawdkopencateorderpull(clt *core.SDKClient, req *wdk.AlibabawdkopencateorderpullAPIRequest, session string) (*wdk.AlibabawdkopencateorderpullAPIResponse, error) {
	var resp wdk.AlibabawdkopencateorderpullAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
