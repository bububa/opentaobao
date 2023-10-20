package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalbasehosstatussync 挂号医院上下线
// alibaba.alihealth.medicalbase.hos.status.sync
//
// 挂号医院上下线
func Alibabaalihealthmedicalbasehosstatussync(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalbasehosstatussyncAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalbasehosstatussyncAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalbasehosstatussyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
