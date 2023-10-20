package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeHouseCooperateBrandUpdate 租房合作品牌更新接口
// alibaba.alihouse.existinghome.house.cooperate.brand.update
//
// 租房合作品牌更新接口
func AlibabaAlihouseExistinghomeHouseCooperateBrandUpdate(clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
