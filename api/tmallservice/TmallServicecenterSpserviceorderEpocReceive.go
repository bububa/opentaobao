package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderEpocReceive 电子保单数据接口
// tmall.servicecenter.spserviceorder.epoc.receive
//
// 电子保单数据回传接口
func TmallServicecenterSpserviceorderEpocReceive(clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderEpocReceiveAPIRequest, session string) (*tmallservice.TmallServicecenterSpserviceorderEpocReceiveAPIResponse, error) {
	var resp tmallservice.TmallServicecenterSpserviceorderEpocReceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
