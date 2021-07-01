package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
售后上门信息同步 API请求
alibaba.ihome.ctom.postsale.onsite.sync

用于三维家同步售后单上门人员和时间信息
*/
type AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest struct {
    model.Params
    // 预约后展示给客户的员工联系方式
    _deliveryPhone   string
    // 提交部门ID（预约的操作人所属部门）
    _unitId   string
    // 备注
    _memo   string
    // 预约事件发生时间
    _time   int64
    // 服务ID
    _serviceId   string
    // 预约目标时间
    _deliveryDate   int64
    // 具体操作人ID（预约人）
    _operatorId   string
    // 联系人员名字
    _deliveryName   string
    // 售后单ID
    _postSalesId   string
    // 配送、安装或上门
    _type   string
    // 三维家补单ID
    _additionalOrderId   string
}

// 初始化AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest对象
func NewAlibabaIhomeCtomPostsaleOnsiteSyncRequest() *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest{
    return &AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.postsale.onsite.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeliveryPhone Setter
// 预约后展示给客户的员工联系方式
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetDeliveryPhone(_deliveryPhone string) error {
    r._deliveryPhone = _deliveryPhone
    r.Set("delivery_phone", _deliveryPhone)
    return nil
}

// DeliveryPhone Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetDeliveryPhone() string {
    return r._deliveryPhone
}
// UnitId Setter
// 提交部门ID（预约的操作人所属部门）
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetUnitId(_unitId string) error {
    r._unitId = _unitId
    r.Set("unit_id", _unitId)
    return nil
}

// UnitId Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetUnitId() string {
    return r._unitId
}
// Memo Setter
// 备注
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetMemo(_memo string) error {
    r._memo = _memo
    r.Set("memo", _memo)
    return nil
}

// Memo Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetMemo() string {
    return r._memo
}
// Time Setter
// 预约事件发生时间
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetTime(_time int64) error {
    r._time = _time
    r.Set("time", _time)
    return nil
}

// Time Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetTime() int64 {
    return r._time
}
// ServiceId Setter
// 服务ID
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetServiceId(_serviceId string) error {
    r._serviceId = _serviceId
    r.Set("service_id", _serviceId)
    return nil
}

// ServiceId Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetServiceId() string {
    return r._serviceId
}
// DeliveryDate Setter
// 预约目标时间
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetDeliveryDate(_deliveryDate int64) error {
    r._deliveryDate = _deliveryDate
    r.Set("delivery_date", _deliveryDate)
    return nil
}

// DeliveryDate Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetDeliveryDate() int64 {
    return r._deliveryDate
}
// OperatorId Setter
// 具体操作人ID（预约人）
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetOperatorId(_operatorId string) error {
    r._operatorId = _operatorId
    r.Set("operator_id", _operatorId)
    return nil
}

// OperatorId Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetOperatorId() string {
    return r._operatorId
}
// DeliveryName Setter
// 联系人员名字
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetDeliveryName(_deliveryName string) error {
    r._deliveryName = _deliveryName
    r.Set("delivery_name", _deliveryName)
    return nil
}

// DeliveryName Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetDeliveryName() string {
    return r._deliveryName
}
// PostSalesId Setter
// 售后单ID
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetPostSalesId(_postSalesId string) error {
    r._postSalesId = _postSalesId
    r.Set("post_sales_id", _postSalesId)
    return nil
}

// PostSalesId Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetPostSalesId() string {
    return r._postSalesId
}
// Type Setter
// 配送、安装或上门
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetType() string {
    return r._type
}
// AdditionalOrderId Setter
// 三维家补单ID
func (r *AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) SetAdditionalOrderId(_additionalOrderId string) error {
    r._additionalOrderId = _additionalOrderId
    r.Set("additional_order_id", _additionalOrderId)
    return nil
}

// AdditionalOrderId Getter
func (r AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest) GetAdditionalOrderId() string {
    return r._additionalOrderId
}
