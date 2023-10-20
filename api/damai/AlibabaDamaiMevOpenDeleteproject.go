package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopendeleteproject 大麦换验平台-第三方对外开放-项目接口deleteProject
// alibaba.damai.mev.open.deleteproject
//
// deleteProject
func Alibabadamaimevopendeleteproject(clt *core.SDKClient, req *damai.AlibabadamaimevopendeleteprojectAPIRequest, session string) (*damai.AlibabadamaimevopendeleteprojectAPIResponse, error) {
	var resp damai.AlibabadamaimevopendeleteprojectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
