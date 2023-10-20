package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingopendatarelationquery 数据关联关系查询
// alibaba.wdk.marketing.open.data.relation.query
//
// 数据关联关系查询
func Alibabawdkmarketingopendatarelationquery(clt *core.SDKClient, req *wdk.AlibabawdkmarketingopendatarelationqueryAPIRequest, session string) (*wdk.AlibabawdkmarketingopendatarelationqueryAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingopendatarelationqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
