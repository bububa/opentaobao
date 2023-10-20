package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousebusinessactivityquery 电商券活动公司数据查询
// alibaba.alihouse.business.activity.query
//
// 电商券活动公司数据查询
func Alibabaalihousebusinessactivityquery(clt *core.SDKClient, req *alihouse.AlibabaalihousebusinessactivityqueryAPIRequest, session string) (*alihouse.AlibabaalihousebusinessactivityqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihousebusinessactivityqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
