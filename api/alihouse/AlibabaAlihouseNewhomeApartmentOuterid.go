package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeApartmentOuterid 公寓更新outerid
// alibaba.alihouse.newhome.apartment.outerid
//
// 公寓更新outerid
func AlibabaAlihouseNewhomeApartmentOuterid(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeApartmentOuteridAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeApartmentOuteridAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
