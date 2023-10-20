package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtstallpayratiosynchronize 同步摊位收银比例
// tmall.nrt.stall.payratio.synchronize
//
// ISV同步摊位收银比例到阿里
func Tmallnrtstallpayratiosynchronize(clt *core.SDKClient, req *nrt.TmallnrtstallpayratiosynchronizeAPIRequest, session string) (*nrt.TmallnrtstallpayratiosynchronizeAPIResponse, error) {
	var resp nrt.TmallnrtstallpayratiosynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
