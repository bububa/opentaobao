package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaocloudprintstdtemplatesget 获取所有的菜鸟标准电子面单模板
// cainiao.cloudprint.stdtemplates.get
//
// 获取菜鸟标准电子面单模板
func Cainiaocloudprintstdtemplatesget(clt *core.SDKClient, req *waybill.CainiaocloudprintstdtemplatesgetAPIRequest, session string) (*waybill.CainiaocloudprintstdtemplatesgetAPIResponse, error) {
	var resp waybill.CainiaocloudprintstdtemplatesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
