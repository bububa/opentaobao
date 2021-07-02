package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// AlibabaAlihealthNrSpuQuery 获取标品库标品信息
// alibaba.alihealth.nr.spu.query
//
// 提供给ERP使用的，获取健康标品库信息
func AlibabaAlihealthNrSpuQuery(clt *core.SDKClient, req *drug.AlibabaAlihealthNrSpuQueryAPIRequest, session string) (*drug.AlibabaAlihealthNrSpuQueryAPIResponse, error) {
	var resp drug.AlibabaAlihealthNrSpuQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
