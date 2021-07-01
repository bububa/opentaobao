package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductGradepriceUpdateAPIRequest
根据sku设置折扣价 API请求
taobao.fenxiao.product.gradeprice.update

供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格 */
type TaobaoFenxiaoProductGradepriceUpdateAPIRequest struct {
	model.Params
	// 交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销)
	_tradeType string
	// 产品Id
	_productId int64
	// skuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选
	_skuId int64
	// 选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如"GRADE,DISTRIBUTOR"
	_targetType string
	// 会员等级的id或者分销商id，例如：”1001,2001,1002”
	_ids []int64
	// 优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25
	_prices []string
}

// New
