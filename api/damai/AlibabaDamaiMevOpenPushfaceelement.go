package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenpushfaceelement 大麦换验平台-第三方对外开放-票面元素接口pushFaceElement
// alibaba.damai.mev.open.pushfaceelement
//
// pushFaceElement
func Alibabadamaimevopenpushfaceelement(clt *core.SDKClient, req *damai.AlibabadamaimevopenpushfaceelementAPIRequest, session string) (*damai.AlibabadamaimevopenpushfaceelementAPIResponse, error) {
	var resp damai.AlibabadamaimevopenpushfaceelementAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
