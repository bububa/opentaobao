package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenpushface 大麦换验平台-第三方对外开放-票面接口pushFace
// alibaba.damai.mev.open.pushface
//
// pushFace
func Alibabadamaimevopenpushface(clt *core.SDKClient, req *damai.AlibabadamaimevopenpushfaceAPIRequest, session string) (*damai.AlibabadamaimevopenpushfaceAPIResponse, error) {
	var resp damai.AlibabadamaimevopenpushfaceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
