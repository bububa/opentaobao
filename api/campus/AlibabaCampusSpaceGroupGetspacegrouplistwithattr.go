package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceGroupGetspacegrouplistwithattr 分页查询空间分组业务属性
// alibaba.campus.space.group.getspacegrouplistwithattr
//
// 分页查询空间分组业务属性
func AlibabaCampusSpaceGroupGetspacegrouplistwithattr(clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest, session string) (*campus.AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse, error) {
	var resp campus.AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
