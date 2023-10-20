package tmallchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaochanneltradeprepayofflineaddAPIRequest 渠道分销供应商上传线下流水预存款（增加） API请求
// taobao.channel.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
type TaobaochanneltradeprepayofflineaddAPIRequest struct {
	model.Params
	// 增加流水
	_offlineAddPrepayParam *TopOfflineAddPrepayDto
}

// NewTaobaochanneltradeprepayofflineaddRequest 初始化TaobaochanneltradeprepayofflineaddAPIRequest对象
func NewTaobaochanneltradeprepayofflineaddRequest() *TaobaochanneltradeprepayofflineaddAPIRequest {
	return &TaobaochanneltradeprepayofflineaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaochanneltradeprepayofflineaddAPIRequest) GetApiMethodName() string {
	return "taobao.channel.trade.prepay.offline.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaochanneltradeprepayofflineaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaochanneltradeprepayofflineaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineAddPrepayParam is OfflineAddPrepayParam Setter
// 增加流水
func (r *TaobaochanneltradeprepayofflineaddAPIRequest) SetOfflineAddPrepayParam(_offlineAddPrepayParam *TopOfflineAddPrepayDto) error {
	r._offlineAddPrepayParam = _offlineAddPrepayParam
	r.Set("offline_add_prepay_param", _offlineAddPrepayParam)
	return nil
}

// GetOfflineAddPrepayParam OfflineAddPrepayParam Getter
func (r TaobaochanneltradeprepayofflineaddAPIRequest) GetOfflineAddPrepayParam() *TopOfflineAddPrepayDto {
	return r._offlineAddPrepayParam
}
