package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelTradeMemoAddAPIRequest
添加一笔交易备注 API请求
taobao.alitrip.travel.trade.memo.add

对一笔交易添加备注 */
type TaobaoAlitripTravelTradeMemoAddAPIRequest struct {
	model.Params
	// 交易ID
	_tid int64
	// 交易备注。最大长度: 1000个字节
	_memo string
	// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
	_flag int64
}

// New
