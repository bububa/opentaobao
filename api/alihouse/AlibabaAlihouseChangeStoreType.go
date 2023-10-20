package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseChangeStoreType 融合店迁移门店
// alibaba.alihouse.change.store.type
//
// 融合店迁移门店
func AlibabaAlihouseChangeStoreType(clt *core.SDKClient, req *alihouse.AlibabaAlihouseChangeStoreTypeAPIRequest, session string) (*alihouse.AlibabaAlihouseChangeStoreTypeAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseChangeStoreTypeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
