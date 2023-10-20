package alihealthcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealthcrm"
)

// Alibabaalihealthalipaypfmdietrecord 用户每日摄入卡路里总量回传接口
// alibaba.alihealth.alipaypfm.diet.record
//
// 用户每日摄入卡路里总量回传接口
func Alibabaalihealthalipaypfmdietrecord(clt *core.SDKClient, req *alihealthcrm.AlibabaalihealthalipaypfmdietrecordAPIRequest, session string) (*alihealthcrm.AlibabaalihealthalipaypfmdietrecordAPIResponse, error) {
	var resp alihealthcrm.AlibabaalihealthalipaypfmdietrecordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
