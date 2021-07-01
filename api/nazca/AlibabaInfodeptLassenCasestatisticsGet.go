package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

/* AlibabaInfodeptLassenCasestatisticsGet
法庭提交和结案案件量接口
alibaba.infodept.lassen.casestatistics.get

功能描述：云嘉为浙江省高院制作数据大屏，需展示网上法庭相关数据，我方为省高院提供浙江省内法院收案和结案的案件量，开放数据接口，供其调取这两组数据。 */
func AlibabaInfodeptLassenCasestatisticsGet(clt *core.SDKClient, req *nazca.AlibabaInfodeptLassenCasestatisticsGetAPIRequest, session string) (*nazca.AlibabaInfodeptLassenCasestatisticsGetAPIResponse, error) {
	var resp nazca.AlibabaInfodeptLassenCasestatisticsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
