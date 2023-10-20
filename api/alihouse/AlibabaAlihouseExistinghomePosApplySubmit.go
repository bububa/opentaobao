package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomePosApplySubmit pos申请接口
// alibaba.alihouse.existinghome.pos.apply.submit
//
// pos申请接口
func AlibabaAlihouseExistinghomePosApplySubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomePosApplySubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomePosApplySubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomePosApplySubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
