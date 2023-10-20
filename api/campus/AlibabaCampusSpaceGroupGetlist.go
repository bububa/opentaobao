package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceGroupGetlist 多条件查询空间分组信息
// alibaba.campus.space.group.getlist
//
// 多条件查询空间分组信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getList
func AlibabaCampusSpaceGroupGetlist(clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetlistAPIRequest, session string) (*campus.AlibabaCampusSpaceGroupGetlistAPIResponse, error) {
	var resp campus.AlibabaCampusSpaceGroupGetlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
