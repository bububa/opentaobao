package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallMsfReceive 签收接口
// tmall.msf.receive
//
// 签收接口
func TmallMsfReceive(clt *core.SDKClient, req *servicecenter.TmallMsfReceiveAPIRequest, resp *servicecenter.TmallMsfReceiveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
