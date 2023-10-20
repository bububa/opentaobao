package waybill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/waybill"
)

// Cainiaocloudprintisvtemplatesget 获取商家使用的标准模板
// cainiao.cloudprint.isvtemplates.get
//
// 获取商家使用的标准模板
func Cainiaocloudprintisvtemplatesget(clt *core.SDKClient, req *waybill.CainiaocloudprintisvtemplatesgetAPIRequest, session string) (*waybill.CainiaocloudprintisvtemplatesgetAPIResponse, error) {
	var resp waybill.CainiaocloudprintisvtemplatesgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
