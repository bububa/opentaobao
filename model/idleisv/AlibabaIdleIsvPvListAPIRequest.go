package idleisv

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvPvListAPIRequest 闲鱼已验货pv查询 API请求
// alibaba.idle.isv.pv.list
//
// 根据闲鱼渠道类目查询对应的品牌和型号清单，供服务商进行选择
type AlibabaIdleIsvPvListAPIRequest struct {
	model.Params
	// 系统自动生成
	_brandModelInfo []IdleNewPubValueDo
	// 闲鱼渠道类目的id
	_channelCatId string
}

// NewAlibabaIdleIsvPvListRequest 初始化AlibabaIdleIsvPvListAPIRequest对象
func NewAlibabaIdleIsvPvListRequest() *AlibabaIdleIsvPvListAPIRequest {
	return &AlibabaIdleIsvPvListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleIsvPvListAPIRequest) Reset() {
	r._brandModelInfo = r._brandModelInfo[:0]
	r._channelCatId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvPvListAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.pv.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvPvListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvPvListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBrandModelInfo is BrandModelInfo Setter
// 系统自动生成
func (r *AlibabaIdleIsvPvListAPIRequest) SetBrandModelInfo(_brandModelInfo []IdleNewPubValueDo) error {
	r._brandModelInfo = _brandModelInfo
	r.Set("brand_model_info", _brandModelInfo)
	return nil
}

// GetBrandModelInfo BrandModelInfo Getter
func (r AlibabaIdleIsvPvListAPIRequest) GetBrandModelInfo() []IdleNewPubValueDo {
	return r._brandModelInfo
}

// SetChannelCatId is ChannelCatId Setter
// 闲鱼渠道类目的id
func (r *AlibabaIdleIsvPvListAPIRequest) SetChannelCatId(_channelCatId string) error {
	r._channelCatId = _channelCatId
	r.Set("channel_cat_id", _channelCatId)
	return nil
}

// GetChannelCatId ChannelCatId Getter
func (r AlibabaIdleIsvPvListAPIRequest) GetChannelCatId() string {
	return r._channelCatId
}

var poolAlibabaIdleIsvPvListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleIsvPvListRequest()
	},
}

// GetAlibabaIdleIsvPvListRequest 从 sync.Pool 获取 AlibabaIdleIsvPvListAPIRequest
func GetAlibabaIdleIsvPvListAPIRequest() *AlibabaIdleIsvPvListAPIRequest {
	return poolAlibabaIdleIsvPvListAPIRequest.Get().(*AlibabaIdleIsvPvListAPIRequest)
}

// ReleaseAlibabaIdleIsvPvListAPIRequest 将 AlibabaIdleIsvPvListAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleIsvPvListAPIRequest(v *AlibabaIdleIsvPvListAPIRequest) {
	v.Reset()
	poolAlibabaIdleIsvPvListAPIRequest.Put(v)
}
