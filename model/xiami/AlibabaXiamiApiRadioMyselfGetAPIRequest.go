package xiami

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiRadioMyselfGetAPIRequest 我的电台 API请求
// alibaba.xiami.api.radio.myself.get
//
// 我的电台
type AlibabaXiamiApiRadioMyselfGetAPIRequest struct {
	model.Params
	// 歌曲数量
	_limit int64
}

// NewAlibabaXiamiApiRadioMyselfGetRequest 初始化AlibabaXiamiApiRadioMyselfGetAPIRequest对象
func NewAlibabaXiamiApiRadioMyselfGetRequest() *AlibabaXiamiApiRadioMyselfGetAPIRequest {
	return &AlibabaXiamiApiRadioMyselfGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaXiamiApiRadioMyselfGetAPIRequest) Reset() {
	r._limit = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaXiamiApiRadioMyselfGetAPIRequest) GetApiMethodName() string {
	return "alibaba.xiami.api.radio.myself.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaXiamiApiRadioMyselfGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaXiamiApiRadioMyselfGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLimit is Limit Setter
// 歌曲数量
func (r *AlibabaXiamiApiRadioMyselfGetAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r AlibabaXiamiApiRadioMyselfGetAPIRequest) GetLimit() int64 {
	return r._limit
}

var poolAlibabaXiamiApiRadioMyselfGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaXiamiApiRadioMyselfGetRequest()
	},
}

// GetAlibabaXiamiApiRadioMyselfGetRequest 从 sync.Pool 获取 AlibabaXiamiApiRadioMyselfGetAPIRequest
func GetAlibabaXiamiApiRadioMyselfGetAPIRequest() *AlibabaXiamiApiRadioMyselfGetAPIRequest {
	return poolAlibabaXiamiApiRadioMyselfGetAPIRequest.Get().(*AlibabaXiamiApiRadioMyselfGetAPIRequest)
}

// ReleaseAlibabaXiamiApiRadioMyselfGetAPIRequest 将 AlibabaXiamiApiRadioMyselfGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaXiamiApiRadioMyselfGetAPIRequest(v *AlibabaXiamiApiRadioMyselfGetAPIRequest) {
	v.Reset()
	poolAlibabaXiamiApiRadioMyselfGetAPIRequest.Put(v)
}
