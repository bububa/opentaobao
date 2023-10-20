package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaifpfulfillwarehousetokenquery 同城令牌打印接口
// alibaba.ifp.fulfill.warehouse.token.query
//
// 用于仓内作业打印包裹信息
func Alibabaifpfulfillwarehousetokenquery(clt *core.SDKClient, req *wdk.AlibabaifpfulfillwarehousetokenqueryAPIRequest, session string) (*wdk.AlibabaifpfulfillwarehousetokenqueryAPIResponse, error) {
	var resp wdk.AlibabaifpfulfillwarehousetokenqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
