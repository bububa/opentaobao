package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderBtobItemQueryAPIRequest 暗拍b2b商品查询 API请求
// alibaba.idle.tender.btob.item.query
//
// 暗拍b2b商品查询
type AlibabaIdleTenderBtobItemQueryAPIRequest struct {
	model.Params
	// 参数
	_param0 *TenderItemListQry
}

// NewAlibabaIdleTenderBtobItemQueryRequest 初始化AlibabaIdleTenderBtobItemQueryAPIRequest对象
func NewAlibabaIdleTenderBtobItemQueryRequest() *AlibabaIdleTenderBtobItemQueryAPIRequest {
	return &AlibabaIdleTenderBtobItemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTenderBtobItemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.btob.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTenderBtobItemQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 参数
func (r *AlibabaIdleTenderBtobItemQueryAPIRequest) SetParam0(_param0 *TenderItemListQry) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaIdleTenderBtobItemQueryAPIRequest) GetParam0() *TenderItemListQry {
	return r._param0
}
