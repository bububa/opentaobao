package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderBtobItemDeleteAPIRequest 暗拍b2b商品下架/删除 API请求
// alibaba.idle.tender.btob.item.delete
//
// 暗拍b2b商品下架/删除
type AlibabaIdleTenderBtobItemDeleteAPIRequest struct {
	model.Params
	// 参数0
	_param0 *TenderItemDeleteCmd
}

// NewAlibabaIdleTenderBtobItemDeleteRequest 初始化AlibabaIdleTenderBtobItemDeleteAPIRequest对象
func NewAlibabaIdleTenderBtobItemDeleteRequest() *AlibabaIdleTenderBtobItemDeleteAPIRequest {
	return &AlibabaIdleTenderBtobItemDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTenderBtobItemDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.btob.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTenderBtobItemDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam0 is Param0 Setter
// 参数0
func (r *AlibabaIdleTenderBtobItemDeleteAPIRequest) SetParam0(_param0 *TenderItemDeleteCmd) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaIdleTenderBtobItemDeleteAPIRequest) GetParam0() *TenderItemDeleteCmd {
	return r._param0
}
