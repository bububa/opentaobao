package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// AlibabaAlihealthPregnancyTaginfoAddrank 点击标签后排序接口
// alibaba.alihealth.pregnancy.taginfo.addrank
//
// 备孕管理--点击标签后排序接口
func AlibabaAlihealthPregnancyTaginfoAddrank(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyTaginfoAddrankAPIRequest, session string) (*alihealthcrm.AlibabaAlihealthPregnancyTaginfoAddrankAPIResponse, error) {
	var resp alihealthcrm.AlibabaAlihealthPregnancyTaginfoAddrankAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
