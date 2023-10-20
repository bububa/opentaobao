package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaocloudprintcustomareaupdate 自定义区内容更新
// cainiao.cloudprint.customarea.update
//
// 自定义区内容更新
func Cainiaocloudprintcustomareaupdate(clt *core.SDKClient, req *waybill.CainiaocloudprintcustomareaupdateAPIRequest, session string) (*waybill.CainiaocloudprintcustomareaupdateAPIResponse, error) {
	var resp waybill.CainiaocloudprintcustomareaupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
