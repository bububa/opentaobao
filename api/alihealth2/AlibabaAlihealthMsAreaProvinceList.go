package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmsareaprovincelist 疫苗预约省份列表查询
// alibaba.alihealth.ms.area.province.list
//
// 微信小程序疫苗预约省份列表查询
func Alibabaalihealthmsareaprovincelist(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmsareaprovincelistAPIRequest, session string) (*alihealth2.AlibabaalihealthmsareaprovincelistAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmsareaprovincelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
