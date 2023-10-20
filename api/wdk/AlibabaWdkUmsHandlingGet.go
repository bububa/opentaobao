package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkumshandlingget 加工单-回流单（新接口）
// alibaba.wdk.ums.handling.get
//
// 加工单-回流单（新接口）
func Alibabawdkumshandlingget(clt *core.SDKClient, req *wdk.AlibabawdkumshandlinggetAPIRequest, session string) (*wdk.AlibabawdkumshandlinggetAPIResponse, error) {
	var resp wdk.AlibabawdkumshandlinggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
