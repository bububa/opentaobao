package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkumsorderget 查询店仓作业单据清单 （库存对账辅助）-回流单
// alibaba.wdk.ums.order.get
//
// 查询店仓作业单据清单 （库存对账辅助）-回流单
func Alibabawdkumsorderget(clt *core.SDKClient, req *wdk.AlibabawdkumsordergetAPIRequest, session string) (*wdk.AlibabawdkumsordergetAPIResponse, error) {
	var resp wdk.AlibabawdkumsordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
