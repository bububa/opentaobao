package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaocloudprintmystdtemplatesget 获取用户使用的菜鸟电子面单模板信息
// cainiao.cloudprint.mystdtemplates.get
//
// 获取用户使用的菜鸟电子面单
func Cainiaocloudprintmystdtemplatesget(clt *core.SDKClient, req *waybill.CainiaocloudprintmystdtemplatesgetAPIRequest, session string) (*waybill.CainiaocloudprintmystdtemplatesgetAPIResponse, error) {
	var resp waybill.CainiaocloudprintmystdtemplatesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
