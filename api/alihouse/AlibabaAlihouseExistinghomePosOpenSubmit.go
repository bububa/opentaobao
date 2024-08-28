package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomePosOpenSubmit pos进件接口
// alibaba.alihouse.existinghome.pos.open.submit
//
// pos进件
func AlibabaAlihouseExistinghomePosOpenSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomePosOpenSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
