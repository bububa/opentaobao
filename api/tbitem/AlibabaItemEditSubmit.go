package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Alibabaitemeditsubmit 商品编辑提交schema信息
// alibaba.item.edit.submit
//
// 商品编辑提交schema信息
func Alibabaitemeditsubmit(clt *core.SDKClient, req *tbitem.AlibabaitemeditsubmitAPIRequest, session string) (*tbitem.AlibabaitemeditsubmitAPIResponse, error) {
	var resp tbitem.AlibabaitemeditsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
