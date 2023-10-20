package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Tmallproductupdateschemaget 产品更新规则获取接口
// tmall.product.update.schema.get
//
// 获取用户更新产品的规则
func Tmallproductupdateschemaget(clt *core.SDKClient, req *tbitem.TmallproductupdateschemagetAPIRequest, session string) (*tbitem.TmallproductupdateschemagetAPIResponse, error) {
	var resp tbitem.TmallproductupdateschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
