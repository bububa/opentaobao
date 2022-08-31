package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeRcSync 阿里房产图文草稿信息同步
// alibaba.alihouse.newhome.rc.sync
//
// 接收图文草稿信息
func AlibabaAlihouseNewhomeRcSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeRcSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeRcSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeRcSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
