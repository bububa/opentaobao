package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemOperateDelete 商品删除
// alibaba.item.operate.delete
//
// 商品删除
func AlibabaItemOperateDelete(clt *core.SDKClient, req *tbitem.AlibabaItemOperateDeleteAPIRequest, session string) (*tbitem.AlibabaItemOperateDeleteAPIResponse, error) {
	var resp tbitem.AlibabaItemOperateDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
