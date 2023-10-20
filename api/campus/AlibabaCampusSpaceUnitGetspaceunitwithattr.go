package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceUnitGetspaceunitwithattr 空间单元id查业务属性实例
// alibaba.campus.space.unit.getspaceunitwithattr
//
// 空间单元id查业务属性实例
func AlibabaCampusSpaceUnitGetspaceunitwithattr(clt *core.SDKClient, req *campus.AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest, resp *campus.AlibabaCampusSpaceUnitGetspaceunitwithattrAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
