package alihealthoutflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthoutflow"
)

// Alibabaalihealthoutflowdrugsupervisionquery 监管平台药品查询
// alibaba.alihealth.outflow.drug.supervision.query
//
// 获取监管平台药品数据
func Alibabaalihealthoutflowdrugsupervisionquery(clt *core.SDKClient, req *alihealthoutflow.AlibabaalihealthoutflowdrugsupervisionqueryAPIRequest, session string) (*alihealthoutflow.AlibabaalihealthoutflowdrugsupervisionqueryAPIResponse, error) {
	var resp alihealthoutflow.AlibabaalihealthoutflowdrugsupervisionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
