package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusadminmapuserlocationinfogetuserlocationinfologs 分时间段获取用户历史位置信息
// alibaba.campus.adminmap.userlocationinfo.getuserlocationinfologs
//
// 分时间段获取用户历史位置信息
func Alibabacampusadminmapuserlocationinfogetuserlocationinfologs(clt *core.SDKClient, req *campus.AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIRequest, session string) (*campus.AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIResponse, error) {
	var resp campus.AlibabacampusadminmapuserlocationinfogetuserlocationinfologsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
