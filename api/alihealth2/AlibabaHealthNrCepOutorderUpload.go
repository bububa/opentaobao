package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabahealthnrcepoutorderupload 线上订单收货验收单、出入库单据生成接口
// alibaba.health.nr.cep.outorder.upload
//
// 线上订单收货验收单、出入库单据生成接口
func Alibabahealthnrcepoutorderupload(clt *core.SDKClient, req *alihealth2.AlibabahealthnrcepoutorderuploadAPIRequest, session string) (*alihealth2.AlibabahealthnrcepoutorderuploadAPIResponse, error) {
	var resp alihealth2.AlibabahealthnrcepoutorderuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
