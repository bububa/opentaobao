package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkorderget 交易订单详情查询
// alibaba.wdk.order.get
//
// 五道口三江单据查询接口
func Alibabawdkorderget(clt *core.SDKClient, req *wdk.AlibabawdkordergetAPIRequest, session string) (*wdk.AlibabawdkordergetAPIResponse, error) {
	var resp wdk.AlibabawdkordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
