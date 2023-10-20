package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomequotationsync 二手房行情数据同步
// alibaba.alihouse.existinghome.quotation.sync
//
// 二手房行情数据同步
func Alibabaalihouseexistinghomequotationsync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomequotationsyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomequotationsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomequotationsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
