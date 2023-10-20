package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacegroupgetspacegrouplistwithattr 分页查询空间分组业务属性
// alibaba.campus.space.group.getspacegrouplistwithattr
//
// 分页查询空间分组业务属性
func Alibabacampusspacegroupgetspacegrouplistwithattr(clt *core.SDKClient, req *campus.AlibabacampusspacegroupgetspacegrouplistwithattrAPIRequest, session string) (*campus.AlibabacampusspacegroupgetspacegrouplistwithattrAPIResponse, error) {
	var resp campus.AlibabacampusspacegroupgetspacegrouplistwithattrAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
