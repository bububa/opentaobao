package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest 通知改期结果 API请求
// alibaba.alihealth.examination.reserve.modify.notify
//
// 体检状态为改期中，服务上通知健康是否改期成功
type AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest struct {
	model.Params
	// 服务商预约凭证
	_uniqReserveCode string
	// 旧的预约日期
	_oldReserveDate string
	// 套餐编码
	_packageCode string
	// 健康预约凭证
	_reserveNumber string
	// 新的预约日期
	_newReserveDate string
	// 商品编码
	_goodsCode string
	// 门店编码
	_storeCode string
	// 拒绝修改的时候需要传递拒绝原因
	_reason string
	// 新的预约时间段开始时间
	_newReserveTimeStart string
	// 新的预约时间段结束时间
	_newReserveTimeEnd string
	// true:同意修改；false:拒绝修改
	_pass bool
}

// NewAlibabaAlihealthExaminationReserveModifyNotifyRequest 初始化AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest对象
func NewAlibabaAlihealthExaminationReserveModifyNotifyRequest() *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest {
	return &AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.modify.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUniqReserveCode is UniqReserveCode Setter
// 服务商预约凭证
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}

// SetOldReserveDate is OldReserveDate Setter
// 旧的预约日期
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetOldReserveDate(_oldReserveDate string) error {
	r._oldReserveDate = _oldReserveDate
	r.Set("old_reserve_date", _oldReserveDate)
	return nil
}

// GetOldReserveDate OldReserveDate Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetOldReserveDate() string {
	return r._oldReserveDate
}

// SetPackageCode is PackageCode Setter
// 套餐编码
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetPackageCode(_packageCode string) error {
	r._packageCode = _packageCode
	r.Set("package_code", _packageCode)
	return nil
}

// GetPackageCode PackageCode Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetPackageCode() string {
	return r._packageCode
}

// SetReserveNumber is ReserveNumber Setter
// 健康预约凭证
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetNewReserveDate is NewReserveDate Setter
// 新的预约日期
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetNewReserveDate(_newReserveDate string) error {
	r._newReserveDate = _newReserveDate
	r.Set("new_reserve_date", _newReserveDate)
	return nil
}

// GetNewReserveDate NewReserveDate Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetNewReserveDate() string {
	return r._newReserveDate
}

// SetGoodsCode is GoodsCode Setter
// 商品编码
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetGoodsCode(_goodsCode string) error {
	r._goodsCode = _goodsCode
	r.Set("goods_code", _goodsCode)
	return nil
}

// GetGoodsCode GoodsCode Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetGoodsCode() string {
	return r._goodsCode
}

// SetStoreCode is StoreCode Setter
// 门店编码
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetReason is Reason Setter
// 拒绝修改的时候需要传递拒绝原因
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetReason() string {
	return r._reason
}

// SetNewReserveTimeStart is NewReserveTimeStart Setter
// 新的预约时间段开始时间
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetNewReserveTimeStart(_newReserveTimeStart string) error {
	r._newReserveTimeStart = _newReserveTimeStart
	r.Set("new_reserve_time_start", _newReserveTimeStart)
	return nil
}

// GetNewReserveTimeStart NewReserveTimeStart Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetNewReserveTimeStart() string {
	return r._newReserveTimeStart
}

// SetNewReserveTimeEnd is NewReserveTimeEnd Setter
// 新的预约时间段结束时间
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetNewReserveTimeEnd(_newReserveTimeEnd string) error {
	r._newReserveTimeEnd = _newReserveTimeEnd
	r.Set("new_reserve_time_end", _newReserveTimeEnd)
	return nil
}

// GetNewReserveTimeEnd NewReserveTimeEnd Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetNewReserveTimeEnd() string {
	return r._newReserveTimeEnd
}

// SetPass is Pass Setter
// true:同意修改；false:拒绝修改
func (r *AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) SetPass(_pass bool) error {
	r._pass = _pass
	r.Set("pass", _pass)
	return nil
}

// GetPass Pass Getter
func (r AlibabaAlihealthExaminationReserveModifyNotifyAPIRequest) GetPass() bool {
	return r._pass
}
