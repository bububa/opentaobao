package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeAgreementSync 二手房电子协议数据同步
// alibaba.alihouse.existinghome.agreement.sync
//
// 二手房电子协议数据同步
func AlibabaAlihouseExistinghomeAgreementSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeAgreementSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeAgreementSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
