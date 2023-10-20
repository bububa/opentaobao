package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkhrworkbenchmokaentryreceiptwrite 摩卡确认入职后往入职单据表写数据接口
// alibaba.wdk.hrworkbench.moka.entry.receipt.write
//
// 摩卡确认入职后往入职单据表写数据接口
func Alibabawdkhrworkbenchmokaentryreceiptwrite(clt *core.SDKClient, req *wdk.AlibabawdkhrworkbenchmokaentryreceiptwriteAPIRequest, session string) (*wdk.AlibabawdkhrworkbenchmokaentryreceiptwriteAPIResponse, error) {
	var resp wdk.AlibabawdkhrworkbenchmokaentryreceiptwriteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
