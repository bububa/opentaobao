package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Tmallservicecenteranomalyrecourseclose 服务投诉问题单关单
// tmall.servicecenter.anomalyrecourse.close
//
// 提供给服务商在投诉单完结时调用，关闭投诉问题工单。
func Tmallservicecenteranomalyrecourseclose(clt *core.SDKClient, req *tmallsc.TmallservicecenteranomalyrecoursecloseAPIRequest, session string) (*tmallsc.TmallservicecenteranomalyrecoursecloseAPIResponse, error) {
	var resp tmallsc.TmallservicecenteranomalyrecoursecloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
