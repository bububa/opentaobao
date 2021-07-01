package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpentradeToolsItemsQueryAPIRequest
交易开放获取商品绑定信息 API请求
taobao.opentrade.tools.items.query

交易开放获取商品绑定信息 */
type TaobaoOpentradeToolsItemsQueryAPIRequest struct {
	model.Params
	// 交易开放C端小程序ID
	_miniappId int64
}

// NewTaobaoOpentradeToolsItemsQueryRequest 初始化TaobaoOpentradeToolsItemsQueryAPIRequest对象
func NewTaobaoOpentradeToolsItemsQueryRequest() *TaobaoOpentradeToolsItemsQueryAPIRequest {
	return &TaobaoOpentradeToolsItemsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeToolsItemsQueryAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.tools.items.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeToolsItemsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MiniappId Setter
// 交易开放C端小程序ID
func (r *TaobaoOpentradeToolsItemsQueryAPIRequest) SetMiniappId(_miniappId int64) error {
	r._miniappId = _miniappId
	r.Set("miniapp_id", _miniappId)
	return nil
}

// Get MiniappId Getter
func (r TaobaoOpentradeToolsItemsQueryAPIRequest) GetMiniappId() int64 {
	return r._miniappId
}
