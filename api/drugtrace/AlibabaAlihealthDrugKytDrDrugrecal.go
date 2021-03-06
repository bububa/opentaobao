package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrDrugrecal 疫苗药品召回
// alibaba.alihealth.drug.kyt.dr.drugrecal
//
// 生产企业发布的召回信息，按照批次进行召回，收货和发货环节的单据处理中调用接口进行查询；
func AlibabaAlihealthDrugKytDrDrugrecal(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrDrugrecalAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrDrugrecalAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrDrugrecalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
