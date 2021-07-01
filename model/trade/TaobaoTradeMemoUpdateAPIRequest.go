package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTradeMemoUpdateAPIRequest
修改交易备注 API请求
taobao.trade.memo.update

需要商家或以上权限才可调用此接口，可重复调用本接口更新交易备注，本接口同时具有添加备注的功能 */
type TaobaoTradeMemoUpdateAPIRequest struct {
	model.Params
	// 交易编号
	_tid int64
	// 卖家交易备注。最大长度: 1000个字节
	_memo string
	// 卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0
	_flag int64
	// 是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值
	_reset bool
}

// New
