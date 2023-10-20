package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaUnitCampusSpaceBookinfoQuery 环路资源信息查询单元环境
// alibaba.unit.campus.space.bookinfo.query
//
// 环路资源信息查询单元环境
func AlibabaUnitCampusSpaceBookinfoQuery(clt *core.SDKClient, req *campus.AlibabaUnitCampusSpaceBookinfoQueryAPIRequest, resp *campus.AlibabaUnitCampusSpaceBookinfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
