package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

/* TaobaoUsergrowthDeliveryProfileReport
标签上报
taobao.usergrowth.delivery.profile.report

渠道上报标签信息 */
func TaobaoUsergrowthDeliveryProfileReport(clt *core.SDKClient, req *usergrowth.TaobaoUsergrowthDeliveryProfileReportAPIRequest, session string) (*usergrowth.TaobaoUsergrowthDeliveryProfileReportAPIResponse, error) {
	var resp usergrowth.TaobaoUsergrowthDeliveryProfileReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
