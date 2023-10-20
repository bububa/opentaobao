package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeSignatureSync 二手房电子签章数据同步
// alibaba.alihouse.existinghome.signature.sync
//
// 二手房电子签章数据同步
func AlibabaAlihouseExistinghomeSignatureSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeSignatureSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeSignatureSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
