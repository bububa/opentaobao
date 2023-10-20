package tmallcampus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcampus"
)

// Tmallcampusauthstatusquery 学生认证状态查询
// tmall.campus.authstatus.query
//
// 学生认证状态查询
func Tmallcampusauthstatusquery(clt *core.SDKClient, req *tmallcampus.TmallcampusauthstatusqueryAPIRequest, session string) (*tmallcampus.TmallcampusauthstatusqueryAPIResponse, error) {
	var resp tmallcampus.TmallcampusauthstatusqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
