package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalscgrowthinteractivesnsconverturl 防封接口
// alibaba.alsc.growth.interactive.sns.converturl
//
// 防封接口
func Alibabaalscgrowthinteractivesnsconverturl(clt *core.SDKClient, req *alsc.AlibabaalscgrowthinteractivesnsconverturlAPIRequest, session string) (*alsc.AlibabaalscgrowthinteractivesnsconverturlAPIResponse, error) {
	var resp alsc.AlibabaalscgrowthinteractivesnsconverturlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
