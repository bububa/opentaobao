package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomePosOpenSubmit pos进件接口
// alibaba.alihouse.existinghome.pos.open.submit
//
// pos进件
func AlibabaAlihouseExistinghomePosOpenSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomePosOpenSubmitAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomePosOpenSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
