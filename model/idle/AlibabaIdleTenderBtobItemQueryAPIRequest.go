package idle

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleTenderBtobItemQueryAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTenderBtobItemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.btob.item.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTenderBtobItemQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTenderBtobItemQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIdleTenderBtobItemQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleTenderBtobItemQueryRequest()
	},
}

// GetAlibabaIdleTenderBtobItemQueryRequest 从 sync.Pool 获取 AlibabaIdleTenderBtobItemQueryAPIRequest
func GetAlibabaIdleTenderBtobItemQueryAPIRequest() *AlibabaIdleTenderBtobItemQueryAPIRequest {
	return poolAlibabaIdleTenderBtobItemQueryAPIRequest.Get().(*AlibabaIdleTenderBtobItemQueryAPIRequest)
}

// ReleaseAlibabaIdleTenderBtobItemQueryAPIRequest 将 AlibabaIdleTenderBtobItemQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleTenderBtobItemQueryAPIRequest(v *AlibabaIdleTenderBtobItemQueryAPIRequest) {
	v.Reset()
	poolAlibabaIdleTenderBtobItemQueryAPIRequest.Put(v)
}
