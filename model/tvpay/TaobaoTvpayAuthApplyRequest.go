package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付申请设备授权 API请求
taobao.tvpay.auth.apply

为用户在指定设备上申请支付授权
*/
type TaobaoTvpayAuthApplyRequest struct {
    model.Params
    // 设备id
    _deviceId   string
    // 请求来源
    _from   string
    // 场景
    _bizScene   string
    // 商品名称
    _itemName   string
    // 授权类型
    _operateType   string
    // 外部订单号
    _outApproveId   string
    // 金额
    _totalFee   string
}

// 初始化TaobaoTvpayAuthApplyRequest对象
func NewTaobaoTvpayAuthApplyRequest() *TaobaoTvpayAuthApplyRequest{
    return &TaobaoTvpayAuthApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayAuthApplyRequest) GetApiMethodName() string {
    return "taobao.tvpay.auth.apply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayAuthApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceId Setter
// 设备id
func (r *TaobaoTvpayAuthApplyRequest) SetDeviceId(_deviceId string) error {
    r._deviceId = _deviceId
    r.Set("device_id", _deviceId)
    return nil
}

// DeviceId Getter
func (r TaobaoTvpayAuthApplyRequest) GetDeviceId() string {
    return r._deviceId
}
// From Setter
// 请求来源
func (r *TaobaoTvpayAuthApplyRequest) SetFrom(_from string) error {
    r._from = _from
    r.Set("from", _from)
    return nil
}

// From Getter
func (r TaobaoTvpayAuthApplyRequest) GetFrom() string {
    return r._from
}
// BizScene Setter
// 场景
func (r *TaobaoTvpayAuthApplyRequest) SetBizScene(_bizScene string) error {
    r._bizScene = _bizScene
    r.Set("biz_scene", _bizScene)
    return nil
}

// BizScene Getter
func (r TaobaoTvpayAuthApplyRequest) GetBizScene() string {
    return r._bizScene
}
// ItemName Setter
// 商品名称
func (r *TaobaoTvpayAuthApplyRequest) SetItemName(_itemName string) error {
    r._itemName = _itemName
    r.Set("item_name", _itemName)
    return nil
}

// ItemName Getter
func (r TaobaoTvpayAuthApplyRequest) GetItemName() string {
    return r._itemName
}
// OperateType Setter
// 授权类型
func (r *TaobaoTvpayAuthApplyRequest) SetOperateType(_operateType string) error {
    r._operateType = _operateType
    r.Set("operate_type", _operateType)
    return nil
}

// OperateType Getter
func (r TaobaoTvpayAuthApplyRequest) GetOperateType() string {
    return r._operateType
}
// OutApproveId Setter
// 外部订单号
func (r *TaobaoTvpayAuthApplyRequest) SetOutApproveId(_outApproveId string) error {
    r._outApproveId = _outApproveId
    r.Set("out_approve_id", _outApproveId)
    return nil
}

// OutApproveId Getter
func (r TaobaoTvpayAuthApplyRequest) GetOutApproveId() string {
    return r._outApproveId
}
// TotalFee Setter
// 金额
func (r *TaobaoTvpayAuthApplyRequest) SetTotalFee(_totalFee string) error {
    r._totalFee = _totalFee
    r.Set("total_fee", _totalFee)
    return nil
}

// TotalFee Getter
func (r TaobaoTvpayAuthApplyRequest) GetTotalFee() string {
    return r._totalFee
}
