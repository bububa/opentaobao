package idle

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleTenderBtobItemDeleteAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTenderBtobItemDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.btob.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTenderBtobItemDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTenderBtobItemDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIdleTenderBtobItemDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleTenderBtobItemDeleteRequest()
	},
}

// GetAlibabaIdleTenderBtobItemDeleteRequest 从 sync.Pool 获取 AlibabaIdleTenderBtobItemDeleteAPIRequest
func GetAlibabaIdleTenderBtobItemDeleteAPIRequest() *AlibabaIdleTenderBtobItemDeleteAPIRequest {
	return poolAlibabaIdleTenderBtobItemDeleteAPIRequest.Get().(*AlibabaIdleTenderBtobItemDeleteAPIRequest)
}

// ReleaseAlibabaIdleTenderBtobItemDeleteAPIRequest 将 AlibabaIdleTenderBtobItemDeleteAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleTenderBtobItemDeleteAPIRequest(v *AlibabaIdleTenderBtobItemDeleteAPIRequest) {
	v.Reset()
	poolAlibabaIdleTenderBtobItemDeleteAPIRequest.Put(v)
}
