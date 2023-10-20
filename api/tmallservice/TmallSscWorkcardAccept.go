package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallsscworkcardaccept 服务商接单完成
// tmall.ssc.workcard.accept
//
// 工单完结
func Tmallsscworkcardaccept(clt *core.SDKClient, req *tmallservice.TmallsscworkcardacceptAPIRequest, session string) (*tmallservice.TmallsscworkcardacceptAPIResponse, error) {
	var resp tmallservice.TmallsscworkcardacceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
