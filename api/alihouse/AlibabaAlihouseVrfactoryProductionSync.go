package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousevrfactoryproductionsync vr生产数据上翻
// alibaba.alihouse.vrfactory.production.sync
//
// vr生产数据上翻
func Alibabaalihousevrfactoryproductionsync(clt *core.SDKClient, req *alihouse.AlibabaalihousevrfactoryproductionsyncAPIRequest, session string) (*alihouse.AlibabaalihousevrfactoryproductionsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousevrfactoryproductionsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
