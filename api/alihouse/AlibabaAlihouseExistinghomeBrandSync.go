package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomebrandsync 二手房公司品牌数据同步
// alibaba.alihouse.existinghome.brand.sync
//
// 二手房公司品牌数据同步
func Alibabaalihouseexistinghomebrandsync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomebrandsyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomebrandsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomebrandsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
