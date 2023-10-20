package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusTopologyGetall 获取管理园区的规则拓扑接口
// alibaba.campus.topology.getall
//
// 获取所属园区的所有规则拓扑图
func AlibabaCampusTopologyGetall(clt *core.SDKClient, req *campus.AlibabaCampusTopologyGetallAPIRequest, resp *campus.AlibabaCampusTopologyGetallAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
