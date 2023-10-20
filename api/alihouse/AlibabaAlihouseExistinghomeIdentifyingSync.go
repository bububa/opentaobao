package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomeidentifyingsync 登陆标识信息同步
// alibaba.alihouse.existinghome.identifying.sync
//
// 登陆标识信息同步
func Alibabaalihouseexistinghomeidentifyingsync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomeidentifyingsyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomeidentifyingsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomeidentifyingsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
