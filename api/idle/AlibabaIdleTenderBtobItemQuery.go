package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleTenderBtobItemQuery 暗拍b2b商品查询
// alibaba.idle.tender.btob.item.query
//
// 暗拍b2b商品查询
func AlibabaIdleTenderBtobItemQuery(clt *core.SDKClient, req *idle.AlibabaIdleTenderBtobItemQueryAPIRequest, session string) (*idle.AlibabaIdleTenderBtobItemQueryAPIResponse, error) {
	var resp idle.AlibabaIdleTenderBtobItemQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
