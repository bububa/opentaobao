package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// Alibabaalihealthpregnancydatasync 四类数据同步
// alibaba.alihealth.pregnancy.data.sync
//
// 经期调整；基础体温；排卵试纸；B超测排数据同步
func Alibabaalihealthpregnancydatasync(clt *core.SDKClient, req *alihealthcrm.AlibabaalihealthpregnancydatasyncAPIRequest, session string) (*alihealthcrm.AlibabaalihealthpregnancydatasyncAPIResponse, error) {
	var resp alihealthcrm.AlibabaalihealthpregnancydatasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
