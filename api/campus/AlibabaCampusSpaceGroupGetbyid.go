package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

/* AlibabaCampusSpaceGroupGetbyid
根据分组ID查询相关的空间分组信息
alibaba.campus.space.group.getbyid

根据分组ID查询相关的空间分组信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getById */
func AlibabaCampusSpaceGroupGetbyid(clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetbyidAPIRequest, session string) (*campus.AlibabaCampusSpaceGroupGetbyidAPIResponse, error) {
	var resp campus.AlibabaCampusSpaceGroupGetbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
