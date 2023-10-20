package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidletenderbtobitemdelete 暗拍b2b商品下架/删除
// alibaba.idle.tender.btob.item.delete
//
// 暗拍b2b商品下架/删除
func Alibabaidletenderbtobitemdelete(clt *core.SDKClient, req *idle.AlibabaidletenderbtobitemdeleteAPIRequest, session string) (*idle.AlibabaidletenderbtobitemdeleteAPIResponse, error) {
	var resp idle.AlibabaidletenderbtobitemdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
