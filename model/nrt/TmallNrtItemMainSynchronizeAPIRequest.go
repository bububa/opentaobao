package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtItemMainSynchronizeAPIRequest
家装新零售主商品同步至阿里 API请求
tmall.nrt.item.main.synchronize

同步红星美凯龙存量商品到阿里 */
type TmallNrtItemMainSynchronizeAPIRequest struct {
	model.Params
	// 摊位id
	_boothId string
	// 叶子类目id
	_cid int64
	// 类目属性
	_props []CategoryPropDto
	// 经销商编码
	_dealerCode string
	// 卖场id
	_mallId string
	// 商家编码
	_outerId string
	// 系统自动生成
	_outerProps *MacallineItemExtDto
	// 价格
	_price string
	// 商品名
	_title string
}

// New
