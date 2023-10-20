package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacetypegetpageresult 分页查询空间类别接口
// alibaba.campus.space.type.getpageresult
//
// 分页查询空间类别接口
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getPageResult
func Alibabacampusspacetypegetpageresult(clt *core.SDKClient, req *campus.AlibabacampusspacetypegetpageresultAPIRequest, session string) (*campus.AlibabacampusspacetypegetpageresultAPIResponse, error) {
	var resp campus.AlibabacampusspacetypegetpageresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
