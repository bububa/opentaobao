package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthmedicalbasethirdevaluatesync 三方评论信息同步
// alibaba.alihealth.medicalbase.third.evaluate.sync
//
// 三方评论信息同步
func Alibabaalihealthmedicalbasethirdevaluatesync(clt *core.SDKClient, req *alihealth2.AlibabaalihealthmedicalbasethirdevaluatesyncAPIRequest, session string) (*alihealth2.AlibabaalihealthmedicalbasethirdevaluatesyncAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthmedicalbasethirdevaluatesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
