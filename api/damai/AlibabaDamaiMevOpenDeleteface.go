package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopendeleteface 大麦换验平台-第三方对外开放-票面接口deleteFace
// alibaba.damai.mev.open.deleteface
//
// deleteFace
func Alibabadamaimevopendeleteface(clt *core.SDKClient, req *damai.AlibabadamaimevopendeletefaceAPIRequest, session string) (*damai.AlibabadamaimevopendeletefaceAPIResponse, error) {
	var resp damai.AlibabadamaimevopendeletefaceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
