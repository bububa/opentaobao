package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceAttrSetattr 新增业务属性实例接口
// alibaba.campus.space.attr.setattr
//
// 新增业务属性实例接口
func AlibabaCampusSpaceAttrSetattr(clt *core.SDKClient, req *campus.AlibabaCampusSpaceAttrSetattrAPIRequest, resp *campus.AlibabaCampusSpaceAttrSetattrAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
