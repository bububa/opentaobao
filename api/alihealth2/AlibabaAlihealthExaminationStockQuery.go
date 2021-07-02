package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthExaminationStockQuery 体检机构对接_体检套餐库存查询
// alibaba.alihealth.examination.stock.query
//
// 体检机构对接_体检套餐库存查询
func AlibabaAlihealthExaminationStockQuery(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthExaminationStockQueryAPIRequest, session string) (*alihealth2.AlibabaAlihealthExaminationStockQueryAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthExaminationStockQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
