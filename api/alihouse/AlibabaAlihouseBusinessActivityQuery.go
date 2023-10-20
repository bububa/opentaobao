package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseBusinessActivityQuery 电商券活动公司数据查询
// alibaba.alihouse.business.activity.query
//
// 电商券活动公司数据查询
func AlibabaAlihouseBusinessActivityQuery(clt *core.SDKClient, req *alihouse.AlibabaAlihouseBusinessActivityQueryAPIRequest, session string) (*alihouse.AlibabaAlihouseBusinessActivityQueryAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseBusinessActivityQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
