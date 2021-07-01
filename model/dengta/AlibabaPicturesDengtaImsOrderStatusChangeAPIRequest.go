package dengta

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天下秀回传订单执行状态变动 API请求
alibaba.pictures.dengta.ims.order.status.change

天下秀回传订单执行状态变动
*/
type AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest struct {
    model.Params
    // 状态发生的时间 2020-01-02 10:02:03
    _changeTime   string
    // 描述，如ims关单，返回关单原因。
    _comments   string
    // 天下秀订单id
    _imsOrderId   string
    // 3=抖音，1-微博 2-微信
    _accountType   int64
    // 扩展字段
    _extJson   string
    // 1:待执行  2:执行中  3:发布  4:完成  5:取消
    _status   int64
}

// 初始化AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest对象
func NewAlibabaPicturesDengtaImsOrderStatusChangeRequest() *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest{
    return &AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.ims.order.status.change"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChangeTime Setter
// 状态发生的时间 2020-01-02 10:02:03
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetChangeTime(_changeTime string) error {
    r._changeTime = _changeTime
    r.Set("change_time", _changeTime)
    return nil
}

// ChangeTime Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetChangeTime() string {
    return r._changeTime
}
// Comments Setter
// 描述，如ims关单，返回关单原因。
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetComments(_comments string) error {
    r._comments = _comments
    r.Set("comments", _comments)
    return nil
}

// Comments Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetComments() string {
    return r._comments
}
// ImsOrderId Setter
// 天下秀订单id
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetImsOrderId(_imsOrderId string) error {
    r._imsOrderId = _imsOrderId
    r.Set("ims_order_id", _imsOrderId)
    return nil
}

// ImsOrderId Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetImsOrderId() string {
    return r._imsOrderId
}
// AccountType Setter
// 3=抖音，1-微博 2-微信
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetAccountType(_accountType int64) error {
    r._accountType = _accountType
    r.Set("account_type", _accountType)
    return nil
}

// AccountType Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetAccountType() int64 {
    return r._accountType
}
// ExtJson Setter
// 扩展字段
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetExtJson(_extJson string) error {
    r._extJson = _extJson
    r.Set("ext_json", _extJson)
    return nil
}

// ExtJson Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetExtJson() string {
    return r._extJson
}
// Status Setter
// 1:待执行  2:执行中  3:发布  4:完成  5:取消
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetStatus() int64 {
    return r._status
}
