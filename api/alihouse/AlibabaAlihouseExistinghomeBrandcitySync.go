package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomebrandcitysync 二手房城市品牌店数据同步
// alibaba.alihouse.existinghome.brandcity.sync
//
// 二手房城市品牌店数据同步
func Alibabaalihouseexistinghomebrandcitysync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomebrandcitysyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomebrandcitysyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomebrandcitysyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
