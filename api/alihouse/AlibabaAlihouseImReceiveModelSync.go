package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseImReceiveModelSync IM承接方式同步
// alibaba.alihouse.im.receive.model.sync
//
// IM承接方式同步
func AlibabaAlihouseImReceiveModelSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseImReceiveModelSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseImReceiveModelSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseImReceiveModelSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
