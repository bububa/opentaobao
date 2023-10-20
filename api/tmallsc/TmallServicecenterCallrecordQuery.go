package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterCallrecordQuery 天猫服务平台服务商查询通话记录接口
// tmall.servicecenter.callrecord.query
//
// 天猫服务平台服务商查询通话记录接口
func TmallServicecenterCallrecordQuery(clt *core.SDKClient, req *tmallsc.TmallServicecenterCallrecordQueryAPIRequest, session string) (*tmallsc.TmallServicecenterCallrecordQueryAPIResponse, error) {
	var resp tmallsc.TmallServicecenterCallrecordQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
