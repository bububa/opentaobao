package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomethreedimensionsync 二手房3D户型信息同步
// alibaba.alihouse.existinghome.threedimension.sync
//
// 二手房3D户型信息同步
func Alibabaalihouseexistinghomethreedimensionsync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomethreedimensionsyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomethreedimensionsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomethreedimensionsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
