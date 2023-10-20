package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseimreceivemodelsync IM承接方式同步
// alibaba.alihouse.im.receive.model.sync
//
// IM承接方式同步
func Alibabaalihouseimreceivemodelsync(clt *core.SDKClient, req *alihouse.AlibabaalihouseimreceivemodelsyncAPIRequest, session string) (*alihouse.AlibabaalihouseimreceivemodelsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseimreceivemodelsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
