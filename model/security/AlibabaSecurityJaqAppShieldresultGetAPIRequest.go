package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqappshieldresultgetAPIRequest 用户查询加固结果 API请求
// alibaba.security.jaq.app.shieldresult.get
//
// 用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包
type AlibabasecurityjaqappshieldresultgetAPIRequest struct {
	model.Params
	// 任务唯一标识
	_itemId string
}

// NewAlibabasecurityjaqappshieldresultgetRequest 初始化AlibabasecurityjaqappshieldresultgetAPIRequest对象
func NewAlibabasecurityjaqappshieldresultgetRequest() *AlibabasecurityjaqappshieldresultgetAPIRequest {
	return &AlibabasecurityjaqappshieldresultgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasecurityjaqappshieldresultgetAPIRequest) GetApiMethodName() string {
	return "alibaba.security.jaq.app.shieldresult.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasecurityjaqappshieldresultgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasecurityjaqappshieldresultgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemId is ItemId Setter
// 任务唯一标识
func (r *AlibabasecurityjaqappshieldresultgetAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabasecurityjaqappshieldresultgetAPIRequest) GetItemId() string {
	return r._itemId
}
