package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelTradeMemoUpdateAPIRequest
修改一笔交易备注 API请求
taobao.alitrip.travel.trade.memo.update

更新一笔交易备注 */
type TaobaoAlitripTravelTradeMemoUpdateAPIRequest struct {
	model.Params
	// 交易ID
	_tid int64
	// 交易备注。最大长度: 1000个字节
	_memo string
	// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
	_flag int64
	// 是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
	_reset bool
}

// New
