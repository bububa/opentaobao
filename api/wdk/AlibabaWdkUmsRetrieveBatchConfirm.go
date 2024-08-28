package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsRetrieveBatchConfirm 批量消息确认
// alibaba.wdk.ums.retrieve.batch.confirm
//
// 批量消息确认
func AlibabaWdkUmsRetrieveBatchConfirm(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkUmsRetrieveBatchConfirmAPIRequest, resp *wdk.AlibabaWdkUmsRetrieveBatchConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
