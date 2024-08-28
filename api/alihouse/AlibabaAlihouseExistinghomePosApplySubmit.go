package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomePosApplySubmit pos申请接口
// alibaba.alihouse.existinghome.pos.apply.submit
//
// pos申请接口
func AlibabaAlihouseExistinghomePosApplySubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomePosApplySubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomePosApplySubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
