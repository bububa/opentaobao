package tmallcampus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcampus"
)

// TmallCampusAuthstatusQuery 学生认证状态查询
// tmall.campus.authstatus.query
//
// 学生认证状态查询
func TmallCampusAuthstatusQuery(clt *core.SDKClient, req *tmallcampus.TmallCampusAuthstatusQueryAPIRequest, session string) (*tmallcampus.TmallCampusAuthstatusQueryAPIResponse, error) {
	var resp tmallcampus.TmallCampusAuthstatusQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
