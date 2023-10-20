package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtstallsynchronize 摊位信息同步
// tmall.nrt.stall.synchronize
//
// 摊位信息同步
func Tmallnrtstallsynchronize(clt *core.SDKClient, req *nrt.TmallnrtstallsynchronizeAPIRequest, session string) (*nrt.TmallnrtstallsynchronizeAPIResponse, error) {
	var resp nrt.TmallnrtstallsynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
