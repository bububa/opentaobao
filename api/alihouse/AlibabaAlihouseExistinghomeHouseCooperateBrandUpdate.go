package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehousecooperatebrandupdate 租房合作品牌更新接口
// alibaba.alihouse.existinghome.house.cooperate.brand.update
//
// 租房合作品牌更新接口
func Alibabaalihouseexistinghomehousecooperatebrandupdate(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehousecooperatebrandupdateAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehousecooperatebrandupdateAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehousecooperatebrandupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
