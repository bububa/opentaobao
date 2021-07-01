package vaccin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/vaccin"
)

/* AlibabaHealthVaccinVaccinateComplete
接种完成反馈接口
alibaba.health.vaccin.vaccinate.complete

ISV 将用户完成接种的疫苗同步给免疫规划中心 */
func AlibabaHealthVaccinVaccinateComplete(clt *core.SDKClient, req *vaccin.AlibabaHealthVaccinVaccinateCompleteAPIRequest, session string) (*vaccin.AlibabaHealthVaccinVaccinateCompleteAPIResponse, error) {
	var resp vaccin.AlibabaHealthVaccinVaccinateCompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
