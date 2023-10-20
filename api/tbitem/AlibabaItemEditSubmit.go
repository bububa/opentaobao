package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemEditSubmit 商品编辑提交schema信息
// alibaba.item.edit.submit
//
// 商品编辑提交schema信息
func AlibabaItemEditSubmit(clt *core.SDKClient, req *tbitem.AlibabaItemEditSubmitAPIRequest, resp *tbitem.AlibabaItemEditSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
