package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillBatchQueryByBatchids 作业小票查询接口
// alibaba.wdk.fulfill.batch.query.by.batchids
//
// 根据节点等条件查询履约单小票信息
func AlibabaWdkFulfillBatchQueryByBatchids(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillBatchQueryByBatchidsAPIRequest, resp *wdk.AlibabaWdkFulfillBatchQueryByBatchidsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
