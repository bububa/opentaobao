package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceTypeGetpageresult 分页查询空间类别接口
// alibaba.campus.space.type.getpageresult
//
// 分页查询空间类别接口
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getPageResult
func AlibabaCampusSpaceTypeGetpageresult(clt *core.SDKClient, req *campus.AlibabaCampusSpaceTypeGetpageresultAPIRequest, session string) (*campus.AlibabaCampusSpaceTypeGetpageresultAPIResponse, error) {
	var resp campus.AlibabaCampusSpaceTypeGetpageresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
