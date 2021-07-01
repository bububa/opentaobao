package caipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCaipiaoGoodsInfoGetAPIRequest
根据卖家id与appkey获取商品信息 API请求
taobao.caipiao.goods.info.get

根据卖家id与appkey获取商品信息。 */
type TaobaoCaipiaoGoodsInfoGetAPIRequest struct {
	model.Params
}

// New
