package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkBmPaiyangStatDataQuery 派样统计数据查询
// alibaba.wdk.bm.paiyang.stat.data.query
//
// 派样统计数据查询
func AlibabaWdkBmPaiyangStatDataQuery(clt *core.SDKClient, req *wdk.AlibabaWdkBmPaiyangStatDataQueryAPIRequest, resp *wdk.AlibabaWdkBmPaiyangStatDataQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
