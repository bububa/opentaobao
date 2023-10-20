package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreservecancelAPIRequest 体检机构对接_预约取消 API请求
// alibaba.alihealth.examination.reserve.cancel
//
// 体检机构对接_体检取消
type AlibabaalihealthexaminationreservecancelAPIRequest struct {
	model.Params
	// 商户唯一码
	_merchantCode string
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 预约时间
	_reserveDate string
	// 体检套餐编码
	_packageCode string
	// 店铺ID
	_storeId string
	// 体检机构预约唯一标识码
	_uniqReserveCode string
}

// NewAlibabaalihealthexaminationreservecancelRequest 初始化AlibabaalihealthexaminationreservecancelAPIRequest对象
func NewAlibabaalihealthexaminationreservecancelRequest() *AlibabaalihealthexaminationreservecancelAPIRequest {
	return &AlibabaalihealthexaminationreservecancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreservecancelAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreservecancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreservecancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMerchantCode is MerchantCode Setter
// 商户唯一码
func (r *AlibabaalihealthexaminationreservecancelAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaalihealthexaminationreservecancelAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetReserveNumber is ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaalihealthexaminationreservecancelAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaalihealthexaminationreservecancelAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetReserveDate is ReserveDate Setter
// 预约时间
func (r *AlibabaalihealthexaminationreservecancelAPIRequest) SetReserveDate(_reserveDate string) error {
	r._reserveDate = _reserveDate
	r.Set("reserve_date", _reserveDate)
	return nil
}

// GetReserveDate ReserveDate Getter
func (r AlibabaalihealthexaminationreservecancelAPIRequest) GetReserveDate() string {
	return r._reserveDate
}

// SetPackageCode is PackageCode Setter
// 体检套餐编码
func (r *AlibabaalihealthexaminationreservecancelAPIRequest) SetPackageCode(_packageCode string) error {
	r._packageCode = _packageCode
	r.Set("package_code", _packageCode)
	return nil
}

// GetPackageCode PackageCode Getter
func (r AlibabaalihealthexaminationreservecancelAPIRequest) GetPackageCode() string {
	return r._packageCode
}

// SetStoreId is StoreId Setter
// 店铺ID
func (r *AlibabaalihealthexaminationreservecancelAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaalihealthexaminationreservecancelAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetUniqReserveCode is UniqReserveCode Setter
// 体检机构预约唯一标识码
func (r *AlibabaalihealthexaminationreservecancelAPIRequest) SetUniqReserveCode(_uniqReserveCode string) error {
	r._uniqReserveCode = _uniqReserveCode
	r.Set("uniq_reserve_code", _uniqReserveCode)
	return nil
}

// GetUniqReserveCode UniqReserveCode Getter
func (r AlibabaalihealthexaminationreservecancelAPIRequest) GetUniqReserveCode() string {
	return r._uniqReserveCode
}
