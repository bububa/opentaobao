package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

// Alibabahealthvaccinnoticeannouncementpublish 支付宝疫苗POV公告通知
// alibaba.health.vaccin.notice.announcement.publish
//
// 支付宝疫苗POV发布公告提醒信息
func Alibabahealthvaccinnoticeannouncementpublish(clt *core.SDKClient, req *vaccin.AlibabahealthvaccinnoticeannouncementpublishAPIRequest, session string) (*vaccin.AlibabahealthvaccinnoticeannouncementpublishAPIResponse, error) {
	var resp vaccin.AlibabahealthvaccinnoticeannouncementpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
