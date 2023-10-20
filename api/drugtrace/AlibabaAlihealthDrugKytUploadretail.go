package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytuploadretail 门店销售扫码回传接口
// alibaba.alihealth.drug.kyt.uploadretail
//
// 门店在销售给顾客时，扫描追溯码的数据按照单据回传；
func Alibabaalihealthdrugkytuploadretail(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytuploadretailAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytuploadretailAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytuploadretailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
