package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletenderbtobitemquery 暗拍b2b商品查询
// alibaba.idle.tender.btob.item.query
//
// 暗拍b2b商品查询
func Alibabaidletenderbtobitemquery(clt *core.SDKClient, req *idle.AlibabaidletenderbtobitemqueryAPIRequest, session string) (*idle.AlibabaidletenderbtobitemqueryAPIResponse, error) {
	var resp idle.AlibabaidletenderbtobitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
