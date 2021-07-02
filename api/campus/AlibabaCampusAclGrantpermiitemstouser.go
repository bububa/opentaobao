package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusAclGrantpermiitemstouser 给人直接授权
// alibaba.campus.acl.grantpermiitemstouser
//
// 给人直接授权
func AlibabaCampusAclGrantpermiitemstouser(clt *core.SDKClient, req *campus.AlibabaCampusAclGrantpermiitemstouserAPIRequest, session string) (*campus.AlibabaCampusAclGrantpermiitemstouserAPIResponse, error) {
	var resp campus.AlibabaCampusAclGrantpermiitemstouserAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
