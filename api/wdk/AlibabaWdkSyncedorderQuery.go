package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdksyncedorderquery 五道口查询同步订单
// alibaba.wdk.syncedorder.query
//
// 外部商户查询同步到五道口的订单
func Alibabawdksyncedorderquery(clt *core.SDKClient, req *wdk.AlibabawdksyncedorderqueryAPIRequest, session string) (*wdk.AlibabawdksyncedorderqueryAPIResponse, error) {
	var resp wdk.AlibabawdksyncedorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
