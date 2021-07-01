package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知改期结果 API请求
alibaba.alihealth.examination.reserve.modify.notify

体检状态为改期中，服务上通知健康是否改期成功
*/
type AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest struct {
    model.Params
    // 旧的预约日期
    _oldReserveDate   string
    // 套餐编码
    _packageCode   string
    // 健康预约凭证
    _reserveNumber   string
    // 新的预约日期
    _newReserveDate   string
    // 服务商预约凭证
    _uniqReserveCode   string
    // 商品编码
    _goodsCode   string
    // 门店编码
    _storeCode   string
    // true:同意修改；false:拒绝修改
    _pass   bool
    // 拒绝修改的时候需要传递拒绝原因
    _reason   string
    // 新的预约时间段开始时间
    _newReserveTimeStart   string
    // 新的预约时间段结束时间
    _newReserveTimeEnd   string
}

// 初始化AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest对象
func NewAlibabaAlihealthExaminationReserveModifyNotifyRequest() *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest{
    return &AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.modify.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OldReserveDate Setter
// 旧的预约日期
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetOldReserveDate(_oldReserveDate string) error {
    r._oldReserveDate = _oldReserveDate
    r.Set("old_reserve_date", _oldReserveDate)
    return nil
}

// OldReserveDate Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetOldReserveDate() string {
    return r._oldReserveDate
}
// PackageCode Setter
// 套餐编码
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetPackageCode(_packageCode string) error {
    r._packageCode = _packageCode
    r.Set("package_code", _packageCode)
    return nil
}

// PackageCode Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetPackageCode() string {
    return r._packageCode
}
// ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetReserveNumber(_reserveNumber string) error {
    r._reserveNumber = _reserveNumber
    r.Set("reserve_number", _reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetReserveNumber() string {
    return r._reserveNumber
}
// NewReserveDate Setter
// 新的预约日期
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetNewReserveDate(_newReserveDate string) error {
    r._newReserveDate = _newReserveDate
    r.Set("new_reserve_date", _newReserveDate)
    return nil
}

// NewReserveDate Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetNewReserveDate() string {
    return r._newReserveDate
}
// UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
    r._uniqReserveCode = _uniqReserveCode
    r.Set("uniq_reserve_code", _uniqReserveCode)
    return nil
}

// UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetUniqReserveCode() string {
    return r._uniqReserveCode
}
// GoodsCode Setter
// 商品编码
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetGoodsCode(_goodsCode string) error {
    r._goodsCode = _goodsCode
    r.Set("goods_code", _goodsCode)
    return nil
}

// GoodsCode Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetGoodsCode() string {
    return r._goodsCode
}
// StoreCode Setter
// 门店编码
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetStoreCode(_storeCode string) error {
    r._storeCode = _storeCode
    r.Set("store_code", _storeCode)
    return nil
}

// StoreCode Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetStoreCode() string {
    return r._storeCode
}
// Pass Setter
// true:同意修改；false:拒绝修改
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetPass(_pass bool) error {
    r._pass = _pass
    r.Set("pass", _pass)
    return nil
}

// Pass Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetPass() bool {
    return r._pass
}
// Reason Setter
// 拒绝修改的时候需要传递拒绝原因
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetReason() string {
    return r._reason
}
// NewReserveTimeStart Setter
// 新的预约时间段开始时间
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetNewReserveTimeStart(_newReserveTimeStart string) error {
    r._newReserveTimeStart = _newReserveTimeStart
    r.Set("new_reserve_time_start", _newReserveTimeStart)
    return nil
}

// NewReserveTimeStart Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetNewReserveTimeStart() string {
    return r._newReserveTimeStart
}
// NewReserveTimeEnd Setter
// 新的预约时间段结束时间
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetNewReserveTimeEnd(_newReserveTimeEnd string) error {
    r._newReserveTimeEnd = _newReserveTimeEnd
    r.Set("new_reserve_time_end", _newReserveTimeEnd)
    return nil
}

// NewReserveTimeEnd Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetNewReserveTimeEnd() string {
    return r._newReserveTimeEnd
}
