package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// Alitriptravelvisaapplicantupdate 飞猪度假-普通签证-申请人进度推进接口
// alitrip.travel.visa.applicant.update
//
// 普通签证订单的申请人进度推进接口，用于商家代用户填写申请人基本信息 或 推进单个申请人的签证进度。
func Alitriptravelvisaapplicantupdate(clt *core.SDKClient, req *traveltrade.AlitriptravelvisaapplicantupdateAPIRequest, session string) (*traveltrade.AlitriptravelvisaapplicantupdateAPIResponse, error) {
	var resp traveltrade.AlitriptravelvisaapplicantupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
