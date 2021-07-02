package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthPregnancyBaseinfoGet 拉取备孕初始化信息
// alibaba.alihealth.pregnancy.baseinfo.get
//
// 拉取备孕初始化信息
func AlibabaAlihealthPregnancyBaseinfoGet(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyBaseinfoGetAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthPregnancyBaseinfoGetAPIResponse, error) {
	var resp alihealthcrm.AlibabaAlihealthPregnancyBaseinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
