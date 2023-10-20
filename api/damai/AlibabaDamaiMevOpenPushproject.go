package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenpushproject 大麦换验平台-第三方对外开放-项目接口pushProject
// alibaba.damai.mev.open.pushproject
//
// pushProject
func Alibabadamaimevopenpushproject(clt *core.SDKClient, req *damai.AlibabadamaimevopenpushprojectAPIRequest, session string) (*damai.AlibabadamaimevopenpushprojectAPIResponse, error) {
	var resp damai.AlibabadamaimevopenpushprojectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
