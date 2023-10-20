package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaocloudprintcustomaresget 获取商家的自定义区模板信息
// cainiao.cloudprint.customares.get
//
// 供isv使用，获取商家的自定义区的模板信息
func Cainiaocloudprintcustomaresget(clt *core.SDKClient, req *waybill.CainiaocloudprintcustomaresgetAPIRequest, session string) (*waybill.CainiaocloudprintcustomaresgetAPIResponse, error) {
	var resp waybill.CainiaocloudprintcustomaresgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
