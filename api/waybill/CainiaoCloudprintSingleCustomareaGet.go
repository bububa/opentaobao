package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaocloudprintsinglecustomareaget 获取商家单一自定义区
// cainiao.cloudprint.single.customarea.get
//
// 商家所有快递公司模板只有一个自定义区
func Cainiaocloudprintsinglecustomareaget(clt *core.SDKClient, req *waybill.CainiaocloudprintsinglecustomareagetAPIRequest, session string) (*waybill.CainiaocloudprintsinglecustomareagetAPIResponse, error) {
	var resp waybill.CainiaocloudprintsinglecustomareagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
