package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthExaminationStockQuery 体检机构对接_体检套餐库存查询
// alibaba.alihealth.examination.stock.query
//
// 体检机构对接_体检套餐库存查询
func AlibabaAlihealthExaminationStockQuery(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthExaminationStockQueryAPIRequest, resp *alihealth2.AlibabaAlihealthExaminationStockQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
