package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeMemoAddAPIRequest
对一笔交易添加备注 API请求
taobao.trade.memo.add

根据登录用户的身份（买家或卖家），自动添加相应的交易备注,不能重复调用些接口添加备注，需要更新备注请用taobao.trade.memo.update */
type TaobaoTradeMemoAddAPIRequest struct {
	model.Params
	// 交易编号
	_tid int64
	// 交易备注。最大长度: 1000个字节
	_memo string
	// 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
	_flag int64
}

// New
