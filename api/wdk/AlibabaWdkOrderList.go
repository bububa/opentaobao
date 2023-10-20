package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkorderlist 五道口订单拉取
// alibaba.wdk.order.list
//
// 五道口交易订单拉取接口
func Alibabawdkorderlist(clt *core.SDKClient, req *wdk.AlibabawdkorderlistAPIRequest, session string) (*wdk.AlibabawdkorderlistAPIResponse, error) {
	var resp wdk.AlibabawdkorderlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
