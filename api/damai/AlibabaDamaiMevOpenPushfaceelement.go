package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenPushfaceelement 大麦换验平台-第三方对外开放-票面元素接口pushFaceElement
// alibaba.damai.mev.open.pushfaceelement
//
// pushFaceElement
func AlibabaDamaiMevOpenPushfaceelement(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenPushfaceelementAPIRequest, session string) (*damai.AlibabaDamaiMevOpenPushfaceelementAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenPushfaceelementAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
