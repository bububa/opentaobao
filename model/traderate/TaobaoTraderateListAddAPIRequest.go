package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTraderateListAddAPIRequest
针对父子订单新增批量评价 API请求
taobao.traderate.list.add

针对父子订单新增批量评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不用再通过该接口进行评价</font>) */
type TaobaoTraderateListAddAPIRequest struct {
	model.Params
	// 交易ID
	_tid int64
	// 评价结果。可选值:good(好评),neutral(中评),bad(差评)
	_result string
	// 评价者角色。可选值:seller(卖家),buyer(买家)
	_role string
	// 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容
	_content string
	// 是否匿名，卖家评不能匿名。可选值:true(匿名),false(非匿名)。 注意：如果买家匿名购买，那么买家的评价默认匿名
	_anony bool
}

// New
