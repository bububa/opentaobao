package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtbranddetailquery 品牌详情查询
// tmall.nrt.branddetail.query
//
// 根据品牌id查询品牌的详细信息
func Tmallnrtbranddetailquery(clt *core.SDKClient, req *nrt.TmallnrtbranddetailqueryAPIRequest, session string) (*nrt.TmallnrtbranddetailqueryAPIResponse, error) {
	var resp nrt.TmallnrtbranddetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
