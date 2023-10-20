package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceGroupGetspacegrouplistwithattr 分页查询空间分组业务属性
// alibaba.campus.space.group.getspacegrouplistwithattr
//
// 分页查询空间分组业务属性
func AlibabaCampusSpaceGroupGetspacegrouplistwithattr(clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest, resp *campus.AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
