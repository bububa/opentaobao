package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomePosApplySubmit pos申请接口
// alibaba.alihouse.existinghome.pos.apply.submit
//
// pos申请接口
func AlibabaAlihouseExistinghomePosApplySubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomePosApplySubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomePosApplySubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
