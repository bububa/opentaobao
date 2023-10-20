package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenteranomalyrecourseremarkupdate 天猫服务平台一键求助单服务商备注更新接口
// tmall.servicecenter.anomalyrecourse.remark.update
//
// 一键求助服务商可以回传备注
func Tmallservicecenteranomalyrecourseremarkupdate(clt *core.SDKClient, req *tmallsc.TmallservicecenteranomalyrecourseremarkupdateAPIRequest, session string) (*tmallsc.TmallservicecenteranomalyrecourseremarkupdateAPIResponse, error) {
	var resp tmallsc.TmallservicecenteranomalyrecourseremarkupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
