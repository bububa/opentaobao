package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeAdviserMessageNotice 催促小B发送短信
// alibaba.alihouse.newhome.adviser.message.notice
//
// 催促小B发送短信
func AlibabaAlihouseNewhomeAdviserMessageNotice(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeAdviserMessageNoticeAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeAdviserMessageNoticeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
