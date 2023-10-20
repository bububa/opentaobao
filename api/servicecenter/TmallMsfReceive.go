package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallmsfreceive 签收接口
// tmall.msf.receive
//
// 签收接口
func Tmallmsfreceive(clt *core.SDKClient, req *servicecenter.TmallmsfreceiveAPIRequest, session string) (*servicecenter.TmallmsfreceiveAPIResponse, error) {
	var resp servicecenter.TmallmsfreceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
