package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseCooperateBrandUpdate 租房合作品牌更新接口
// alibaba.alihouse.existinghome.house.cooperate.brand.update
//
// 租房合作品牌更新接口
func AlibabaAlihouseExistinghomeHouseCooperateBrandUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest, session string) (*alihouse.AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
