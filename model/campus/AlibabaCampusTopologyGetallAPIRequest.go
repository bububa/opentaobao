package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusTopologyGetallAPIRequest
获取管理园区的规则拓扑接口 API请求
alibaba.campus.topology.getall

获取所属园区的所有规则拓扑图 */
type AlibabaCampusTopologyGetallAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
	// 公司id
	_companyId int64
	// 系统id
	_systemId string
}

// New
