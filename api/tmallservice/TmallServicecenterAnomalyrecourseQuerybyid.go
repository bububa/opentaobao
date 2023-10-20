package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenteranomalyrecoursequerybyid 根据一键求助id查询指定服务商的一键求助单
// tmall.servicecenter.anomalyrecourse.querybyid
//
// 根据一键求助id查询指定服务商的一键求助单
func Tmallservicecenteranomalyrecoursequerybyid(clt *core.SDKClient, req *tmallservice.TmallservicecenteranomalyrecoursequerybyidAPIRequest, session string) (*tmallservice.TmallservicecenteranomalyrecoursequerybyidAPIResponse, error) {
	var resp tmallservice.TmallservicecenteranomalyrecoursequerybyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
