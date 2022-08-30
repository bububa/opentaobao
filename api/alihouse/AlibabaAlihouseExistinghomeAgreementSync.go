package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeAgreementSync 二手房电子协议数据同步
// alibaba.alihouse.existinghome.agreement.sync
//
// 二手房电子协议数据同步
func AlibabaAlihouseExistinghomeAgreementSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeAgreementSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeAgreementSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeAgreementSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
