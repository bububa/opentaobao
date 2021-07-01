package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTraderateImprImprwordsGetAPIRequest
评价大家印象印象短语接口 API请求
taobao.traderate.impr.imprwords.get

根据淘宝后台类目的一级类目和叶子类目 */
type TaobaoTraderateImprImprwordsGetAPIRequest struct {
	model.Params
	// 淘宝一级类目id
	_catRootId int64
	// 淘宝叶子类目id
	_catLeafId int64
}

// New
