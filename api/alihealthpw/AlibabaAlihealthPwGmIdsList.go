package alihealthpw

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthpw"
)

// AlibabaAlihealthPwGmIdsList 同情用药根据申请单列表查询申请单
// alibaba.alihealth.pw.gm.ids.list
//
// 同情用药根据申请单列表查询申请单
func AlibabaAlihealthPwGmIdsList(ctx context.Context, clt *core.SDKClient, req *alihealthpw.AlibabaAlihealthPwGmIdsListAPIRequest, resp *alihealthpw.AlibabaAlihealthPwGmIdsListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
