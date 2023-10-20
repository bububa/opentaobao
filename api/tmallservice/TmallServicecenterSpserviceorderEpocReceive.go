package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterspserviceorderepocreceive 电子保单数据接口
// tmall.servicecenter.spserviceorder.epoc.receive
//
// 电子保单数据回传接口
func Tmallservicecenterspserviceorderepocreceive(clt *core.SDKClient, req *tmallservice.TmallservicecenterspserviceorderepocreceiveAPIRequest, session string) (*tmallservice.TmallservicecenterspserviceorderepocreceiveAPIResponse, error) {
	var resp tmallservice.TmallservicecenterspserviceorderepocreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
