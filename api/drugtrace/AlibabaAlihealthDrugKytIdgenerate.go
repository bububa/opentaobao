package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytidgenerate 终端(医疗机构|零售药店)ID生成接口
// alibaba.alihealth.drug.kyt.idgenerate
//
// 终端(医疗机构|零售药店)ID生成接口
func Alibabaalihealthdrugkytidgenerate(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytidgenerateAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytidgenerateAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytidgenerateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
