package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalbasedeptstatussync 挂号科室上下线
// alibaba.alihealth.medicalbase.dept.status.sync
//
// 挂号医院上下线
func Alibabaalihealthmedicalbasedeptstatussync(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalbasedeptstatussyncAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalbasedeptstatussyncAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalbasedeptstatussyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
