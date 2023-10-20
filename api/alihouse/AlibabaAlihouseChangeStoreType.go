package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseChangeStoreType 融合店迁移门店
// alibaba.alihouse.change.store.type
//
// 融合店迁移门店
func AlibabaAlihouseChangeStoreType(clt *core.SDKClient, req *alihouse.AlibabaAlihouseChangeStoreTypeAPIRequest, resp *alihouse.AlibabaAlihouseChangeStoreTypeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
