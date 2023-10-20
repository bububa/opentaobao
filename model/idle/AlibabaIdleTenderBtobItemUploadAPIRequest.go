package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderBtobItemUploadAPIRequest 暗拍发布/编辑B2B商品 API请求
// alibaba.idle.tender.btob.item.upload
//
// 暗拍发布/编辑B2B商品
type AlibabaIdleTenderBtobItemUploadAPIRequest struct {
	model.Params
	// 参数
	_param0 *TenderItemUploadCmd
}

// NewAlibabaIdleTenderBtobItemUploadRequest 初始化AlibabaIdleTenderBtobItemUploadAPIRequest对象
func NewAlibabaIdleTenderBtobItemUploadRequest() *AlibabaIdleTenderBtobItemUploadAPIRequest {
	return &AlibabaIdleTenderBtobItemUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleTenderBtobItemUploadAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleTenderBtobItemUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.tender.btob.item.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleTenderBtobItemUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleTenderBtobItemUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 参数
func (r *AlibabaIdleTenderBtobItemUploadAPIRequest) SetParam0(_param0 *TenderItemUploadCmd) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaIdleTenderBtobItemUploadAPIRequest) GetParam0() *TenderItemUploadCmd {
	return r._param0
}

var poolAlibabaIdleTenderBtobItemUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleTenderBtobItemUploadRequest()
	},
}

// GetAlibabaIdleTenderBtobItemUploadRequest 从 sync.Pool 获取 AlibabaIdleTenderBtobItemUploadAPIRequest
func GetAlibabaIdleTenderBtobItemUploadAPIRequest() *AlibabaIdleTenderBtobItemUploadAPIRequest {
	return poolAlibabaIdleTenderBtobItemUploadAPIRequest.Get().(*AlibabaIdleTenderBtobItemUploadAPIRequest)
}

// ReleaseAlibabaIdleTenderBtobItemUploadAPIRequest 将 AlibabaIdleTenderBtobItemUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleTenderBtobItemUploadAPIRequest(v *AlibabaIdleTenderBtobItemUploadAPIRequest) {
	v.Reset()
	poolAlibabaIdleTenderBtobItemUploadAPIRequest.Put(v)
}
