package tvpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayAuthApplyAPIRequest tv支付申请设备授权 API请求
// taobao.tvpay.auth.apply
//
// 为用户在指定设备上申请支付授权
type TaobaoTvpayAuthApplyAPIRequest struct {
	model.Params
	// 设备id
	_deviceId string
	// 请求来源
	_from string
	// 场景
	_bizScene string
	// 商品名称
	_itemName string
	// 授权类型
	_operateType string
	// 外部订单号
	_outApproveId string
	// 金额
	_totalFee string
}

// NewTaobaoTvpayAuthApplyRequest 初始化TaobaoTvpayAuthApplyAPIRequest对象
func NewTaobaoTvpayAuthApplyRequest() *TaobaoTvpayAuthApplyAPIRequest {
	return &TaobaoTvpayAuthApplyAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTvpayAuthApplyAPIRequest) Reset() {
	r._deviceId = ""
	r._from = ""
	r._bizScene = ""
	r._itemName = ""
	r._operateType = ""
	r._outApproveId = ""
	r._totalFee = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAuthApplyAPIRequest) GetApiMethodName() string {
	return "taobao.tvpay.auth.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAuthApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTvpayAuthApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceId is DeviceId Setter
// 设备id
func (r *TaobaoTvpayAuthApplyAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r TaobaoTvpayAuthApplyAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetFrom is From Setter
// 请求来源
func (r *TaobaoTvpayAuthApplyAPIRequest) SetFrom(_from string) error {
	r._from = _from
	r.Set("from", _from)
	return nil
}

// GetFrom From Getter
func (r TaobaoTvpayAuthApplyAPIRequest) GetFrom() string {
	return r._from
}

// SetBizScene is BizScene Setter
// 场景
func (r *TaobaoTvpayAuthApplyAPIRequest) SetBizScene(_bizScene string) error {
	r._bizScene = _bizScene
	r.Set("biz_scene", _bizScene)
	return nil
}

// GetBizScene BizScene Getter
func (r TaobaoTvpayAuthApplyAPIRequest) GetBizScene() string {
	return r._bizScene
}

// SetItemName is ItemName Setter
// 商品名称
func (r *TaobaoTvpayAuthApplyAPIRequest) SetItemName(_itemName string) error {
	r._itemName = _itemName
	r.Set("item_name", _itemName)
	return nil
}

// GetItemName ItemName Getter
func (r TaobaoTvpayAuthApplyAPIRequest) GetItemName() string {
	return r._itemName
}

// SetOperateType is OperateType Setter
// 授权类型
func (r *TaobaoTvpayAuthApplyAPIRequest) SetOperateType(_operateType string) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// GetOperateType OperateType Getter
func (r TaobaoTvpayAuthApplyAPIRequest) GetOperateType() string {
	return r._operateType
}

// SetOutApproveId is OutApproveId Setter
// 外部订单号
func (r *TaobaoTvpayAuthApplyAPIRequest) SetOutApproveId(_outApproveId string) error {
	r._outApproveId = _outApproveId
	r.Set("out_approve_id", _outApproveId)
	return nil
}

// GetOutApproveId OutApproveId Getter
func (r TaobaoTvpayAuthApplyAPIRequest) GetOutApproveId() string {
	return r._outApproveId
}

// SetTotalFee is TotalFee Setter
// 金额
func (r *TaobaoTvpayAuthApplyAPIRequest) SetTotalFee(_totalFee string) error {
	r._totalFee = _totalFee
	r.Set("total_fee", _totalFee)
	return nil
}

// GetTotalFee TotalFee Getter
func (r TaobaoTvpayAuthApplyAPIRequest) GetTotalFee() string {
	return r._totalFee
}

var poolTaobaoTvpayAuthApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTvpayAuthApplyRequest()
	},
}

// GetTaobaoTvpayAuthApplyRequest 从 sync.Pool 获取 TaobaoTvpayAuthApplyAPIRequest
func GetTaobaoTvpayAuthApplyAPIRequest() *TaobaoTvpayAuthApplyAPIRequest {
	return poolTaobaoTvpayAuthApplyAPIRequest.Get().(*TaobaoTvpayAuthApplyAPIRequest)
}

// ReleaseTaobaoTvpayAuthApplyAPIRequest 将 TaobaoTvpayAuthApplyAPIRequest 放入 sync.Pool
func ReleaseTaobaoTvpayAuthApplyAPIRequest(v *TaobaoTvpayAuthApplyAPIRequest) {
	v.Reset()
	poolTaobaoTvpayAuthApplyAPIRequest.Put(v)
}
