package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaotradeprepayofflineaddAPIRequest 线下预存款流水增加 API请求
// taobao.fenxiao.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
type TaobaofenxiaotradeprepayofflineaddAPIRequest struct {
	model.Params
	// 增加流水
	_offlineAddPrepayParam *TopOfflineAddPrepayDto
}

// NewTaobaofenxiaotradeprepayofflineaddRequest 初始化TaobaofenxiaotradeprepayofflineaddAPIRequest对象
func NewTaobaofenxiaotradeprepayofflineaddRequest() *TaobaofenxiaotradeprepayofflineaddAPIRequest {
	return &TaobaofenxiaotradeprepayofflineaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaotradeprepayofflineaddAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.trade.prepay.offline.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaotradeprepayofflineaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaotradeprepayofflineaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineAddPrepayParam is OfflineAddPrepayParam Setter
// 增加流水
func (r *TaobaofenxiaotradeprepayofflineaddAPIRequest) SetOfflineAddPrepayParam(_offlineAddPrepayParam *TopOfflineAddPrepayDto) error {
	r._offlineAddPrepayParam = _offlineAddPrepayParam
	r.Set("offline_add_prepay_param", _offlineAddPrepayParam)
	return nil
}

// GetOfflineAddPrepayParam OfflineAddPrepayParam Getter
func (r TaobaofenxiaotradeprepayofflineaddAPIRequest) GetOfflineAddPrepayParam() *TopOfflineAddPrepayDto {
	return r._offlineAddPrepayParam
}
