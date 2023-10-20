package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// Taobaofiveecompanyget 查询商信息
// taobao.fivee.company.get
//
// 资质共享平台查询商信息
func Taobaofiveecompanyget(clt *core.SDKClient, req *fivee.TaobaofiveecompanygetAPIRequest, session string) (*fivee.TaobaofiveecompanygetAPIResponse, error) {
	var resp fivee.TaobaofiveecompanygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
