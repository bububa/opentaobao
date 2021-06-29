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
type AlibabaPicturesDengtaImsOrderStatusChangeRequest struct {
    model.Params
    // 状态发生的时间 2020-01-02 10:02:03
    changeTime   string
    // 描述，如ims关单，返回关单原因。
    comments   string
    // 天下秀订单id
    imsOrderId   string
    // 3=抖音，1-微博 2-微信
    accountType   int64
    // 扩展字段
    extJson   string
    // 1:待执行  2:执行中  3:发布  4:完成  5:取消
    status   int64
}

// 初始化AlibabaPicturesDengtaImsOrderStatusChangeRequest对象
func NewAlibabaPicturesDengtaImsOrderStatusChangeRequest() *AlibabaPicturesDengtaImsOrderStatusChangeRequest{
    return &AlibabaPicturesDengtaImsOrderStatusChangeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaImsOrderStatusChangeRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.ims.order.status.change"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaImsOrderStatusChangeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChangeTime Setter
// 状态发生的时间 2020-01-02 10:02:03
func (r *AlibabaPicturesDengtaImsOrderStatusChangeRequest) SetChangeTime(changeTime string) error {
    r.changeTime = changeTime
    r.Set("change_time", changeTime)
    return nil
}

// ChangeTime Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeRequest) GetChangeTime() string {
    return r.changeTime
}
// Comments Setter
// 描述，如ims关单，返回关单原因。
func (r *AlibabaPicturesDengtaImsOrderStatusChangeRequest) SetComments(comments string) error {
    r.comments = comments
    r.Set("comments", comments)
    return nil
}

// Comments Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeRequest) GetComments() string {
    return r.comments
}
// ImsOrderId Setter
// 天下秀订单id
func (r *AlibabaPicturesDengtaImsOrderStatusChangeRequest) SetImsOrderId(imsOrderId string) error {
    r.imsOrderId = imsOrderId
    r.Set("ims_order_id", imsOrderId)
    return nil
}

// ImsOrderId Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeRequest) GetImsOrderId() string {
    return r.imsOrderId
}
// AccountType Setter
// 3=抖音，1-微博 2-微信
func (r *AlibabaPicturesDengtaImsOrderStatusChangeRequest) SetAccountType(accountType int64) error {
    r.accountType = accountType
    r.Set("account_type", accountType)
    return nil
}

// AccountType Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeRequest) GetAccountType() int64 {
    return r.accountType
}
// ExtJson Setter
// 扩展字段
func (r *AlibabaPicturesDengtaImsOrderStatusChangeRequest) SetExtJson(extJson string) error {
    r.extJson = extJson
    r.Set("ext_json", extJson)
    return nil
}

// ExtJson Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeRequest) GetExtJson() string {
    return r.extJson
}
// Status Setter
// 1:待执行  2:执行中  3:发布  4:完成  5:取消
func (r *AlibabaPicturesDengtaImsOrderStatusChangeRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeRequest) GetStatus() int64 {
    return r.status
}
