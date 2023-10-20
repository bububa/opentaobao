package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBankcardSync 结算账号同步
// alibaba.alihouse.existinghome.bankcard.sync
//
// 结算账号同步
func AlibabaAlihouseExistinghomeBankcardSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBankcardSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeBankcardSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeBankcardSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
