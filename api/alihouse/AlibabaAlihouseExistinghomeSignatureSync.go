package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeSignatureSync 二手房电子签章数据同步
// alibaba.alihouse.existinghome.signature.sync
//
// 二手房电子签章数据同步
func AlibabaAlihouseExistinghomeSignatureSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeSignatureSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeSignatureSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeSignatureSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
