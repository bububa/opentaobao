package customizemarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/customizemarket"
)

// Alibabataobaoindustrypetprofilequery 用户宠物列表查询
// alibaba.taobao.industry.pet.profile.query
//
// 用户宠物列表查询
func Alibabataobaoindustrypetprofilequery(clt *core.SDKClient, req *customizemarket.AlibabataobaoindustrypetprofilequeryAPIRequest, session string) (*customizemarket.AlibabataobaoindustrypetprofilequeryAPIResponse, error) {
	var resp customizemarket.AlibabataobaoindustrypetprofilequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
