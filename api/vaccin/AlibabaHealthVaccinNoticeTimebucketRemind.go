package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinnoticetimebucketremind 疫苗预约时间段提醒
// alibaba.health.vaccin.notice.timebucket.remind
//
// 疫苗预约时间段提醒
func Alibabahealthvaccinnoticetimebucketremind(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinnoticetimebucketremindAPIRequest, session string) (*vaccin.AlibabahealthvaccinnoticetimebucketremindAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinnoticetimebucketremindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
