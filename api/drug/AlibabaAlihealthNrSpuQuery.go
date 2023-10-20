package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// Alibabaalihealthnrspuquery 获取标品库标品信息
// alibaba.alihealth.nr.spu.query
//
// 提供给ERP使用的，获取健康标品库信息
func Alibabaalihealthnrspuquery(clt *core.SDKClient, req *drug.AlibabaalihealthnrspuqueryAPIRequest, session string) (*drug.AlibabaalihealthnrspuqueryAPIResponse, error) {
	var resp drug.AlibabaalihealthnrspuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
