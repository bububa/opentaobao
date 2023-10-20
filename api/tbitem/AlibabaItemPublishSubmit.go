package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Alibabaitempublishsubmit 商品发布
// alibaba.item.publish.submit
//
// 新商品发布，提交商品发布信息
func Alibabaitempublishsubmit(clt *core.SDKClient, req *tbitem.AlibabaitempublishsubmitAPIRequest, session string) (*tbitem.AlibabaitempublishsubmitAPIResponse, error) {
	var resp tbitem.AlibabaitempublishsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
