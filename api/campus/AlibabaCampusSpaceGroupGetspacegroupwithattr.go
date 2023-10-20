package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceGroupGetspacegroupwithattr 空间分组id查业务属性实例
// alibaba.campus.space.group.getspacegroupwithattr
//
// 空间分组id查业务属性实例
func AlibabaCampusSpaceGroupGetspacegroupwithattr(clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest, resp *campus.AlibabaCampusSpaceGroupGetspacegroupwithattrAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
