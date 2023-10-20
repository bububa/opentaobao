package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopendeletefaceelement 大麦换验平台-第三方对外开放-票面元素接口deleteFaceElement
// alibaba.damai.mev.open.delete.faceelement
//
// deleteFaceElement
func Alibabadamaimevopendeletefaceelement(clt *core.SDKClient, req *damai.AlibabadamaimevopendeletefaceelementAPIRequest, session string) (*damai.AlibabadamaimevopendeletefaceelementAPIResponse, error) {
	var resp damai.AlibabadamaimevopendeletefaceelementAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
