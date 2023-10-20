package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusaclgrantpermiitemstouser 给人直接授权
// alibaba.campus.acl.grantpermiitemstouser
//
// 给人直接授权
func Alibabacampusaclgrantpermiitemstouser(clt *core.SDKClient, req *campus.AlibabacampusaclgrantpermiitemstouserAPIRequest, session string) (*campus.AlibabacampusaclgrantpermiitemstouserAPIResponse, error) {
	var resp campus.AlibabacampusaclgrantpermiitemstouserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
