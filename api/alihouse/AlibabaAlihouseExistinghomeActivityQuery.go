package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomeactivityquery 五福活动经纪人获取
// alibaba.alihouse.existinghome.activity.query
//
// 五福活动经纪人获取
func Alibabaalihouseexistinghomeactivityquery(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomeactivityqueryAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomeactivityqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomeactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
