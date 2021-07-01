package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

/* AlibabaDamaiMevOpenDeleteproject
大麦换验平台-第三方对外开放-项目接口deleteProject
alibaba.damai.mev.open.deleteproject

deleteProject */
func AlibabaDamaiMevOpenDeleteproject(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenDeleteprojectAPIRequest, session string) (*damai.AlibabaDamaiMevOpenDeleteprojectAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenDeleteprojectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
