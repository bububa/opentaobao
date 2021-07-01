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

// NewTaobaoTraderateImprImprwordsGetRequest 初始化TaobaoTraderateImprImprwordsGetAPIRequest对象
func NewTaobaoTraderateImprImprwordsGetRequest() *TaobaoTraderateImprImprwordsGetAPIRequest {
	return &TaobaoTraderateImprImprwordsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTraderateImprImprwordsGetAPIRequest) GetApiMethodName() string {
	return "taobao.traderate.impr.imprwords.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTraderateImprImprwordsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CatRootId Setter
// 淘宝一级类目id
func (r *TaobaoTraderateImprImprwordsGetAPIRequest) SetCatRootId(_catRootId int64) error {
	r._catRootId = _catRootId
	r.Set("cat_root_id", _catRootId)
	return nil
}

// Get CatRootId Getter
func (r TaobaoTraderateImprImprwordsGetAPIRequest) GetCatRootId() int64 {
	return r._catRootId
}

// Set is CatLeafId Setter
// 淘宝叶子类目id
func (r *TaobaoTraderateImprImprwordsGetAPIRequest) SetCatLeafId(_catLeafId int64) error {
	r._catLeafId = _catLeafId
	r.Set("cat_leaf_id", _catLeafId)
	return nil
}

// Get CatLeafId Getter
func (r TaobaoTraderateImprImprwordsGetAPIRequest) GetCatLeafId() int64 {
	return r._catLeafId
}
