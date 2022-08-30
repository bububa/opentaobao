package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeApartmentOuterid 公寓更新outerid
// alibaba.alihouse.newhome.apartment.outerid
//
// 公寓更新outerid
func AlibabaAlihouseNewhomeApartmentOuterid(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeApartmentOuteridAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeApartmentOuteridAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeApartmentOuteridAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
