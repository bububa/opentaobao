package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfillbatchquerybybatchids 作业小票查询接口
// alibaba.wdk.fulfill.batch.query.by.batchids
//
// 根据节点等条件查询履约单小票信息
func Alibabawdkfulfillbatchquerybybatchids(clt *core.SDKClient, req *wdk.AlibabawdkfulfillbatchquerybybatchidsAPIRequest, session string) (*wdk.AlibabawdkfulfillbatchquerybybatchidsAPIResponse, error) {
	var resp wdk.AlibabawdkfulfillbatchquerybybatchidsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
