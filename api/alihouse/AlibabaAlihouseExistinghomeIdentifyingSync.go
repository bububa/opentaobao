package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeIdentifyingSync 登陆标识信息同步
// alibaba.alihouse.existinghome.identifying.sync
//
// 登陆标识信息同步
func AlibabaAlihouseExistinghomeIdentifyingSync(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeIdentifyingSyncAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeIdentifyingSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
