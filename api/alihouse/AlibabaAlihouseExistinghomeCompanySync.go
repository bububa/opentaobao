package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomecompanysync 二手房公司同步接口
// alibaba.alihouse.existinghome.company.sync
//
// 二手房公司同步接口
func Alibabaalihouseexistinghomecompanysync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomecompanysyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomecompanysyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomecompanysyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
