package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// Alibabaailabuserauthorizedquery 查询授权状态接口
// alibaba.ailab.user.authorized.query
//
// 查询三方用户授权状态
func Alibabaailabuserauthorizedquery(clt *core.SDKClient, req *alilabs.AlibabaailabuserauthorizedqueryAPIRequest, session string) (*alilabs.AlibabaailabuserauthorizedqueryAPIResponse, error) {
	var resp alilabs.AlibabaailabuserauthorizedqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
