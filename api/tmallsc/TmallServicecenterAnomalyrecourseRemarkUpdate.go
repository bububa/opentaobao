package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseRemarkUpdate 天猫服务平台一键求助单服务商备注更新接口
// tmall.servicecenter.anomalyrecourse.remark.update
//
// 一键求助服务商可以回传备注
func TmallServicecenterAnomalyrecourseRemarkUpdate(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseRemarkUpdateAPIRequest, session string) (*tmallsc.TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse, error) {
	var resp tmallsc.TmallServicecenterAnomalyrecourseRemarkUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
