package xhotel

import (
	"sync"
)

// DsNhotelInfoDo 结构体
type DsNhotelInfoDo struct {
	// onlineDate
	OnlineDate string `json:"online_date,omitempty" xml:"online_date,omitempty"`
	// signDate
	SignDate string `json:"sign_date,omitempty" xml:"sign_date,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// outerId
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// canSellDate
	CanSellDate string `json:"can_sell_date,omitempty" xml:"can_sell_date,omitempty"`
	// onlineStatus
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
	// qualificationType
	QualificationType int64 `json:"qualification_type,omitempty" xml:"qualification_type,omitempty"`
	// accountType
	AccountType int64 `json:"account_type,omitempty" xml:"account_type,omitempty"`
	// trainStatus
	TrainStatus int64 `json:"train_status,omitempty" xml:"train_status,omitempty"`
	// commLetterStatus
	CommLetterStatus int64 `json:"comm_letter_status,omitempty" xml:"comm_letter_status,omitempty"`
	// privateReceiptAccountStatus
	PrivateReceiptAccountStatus int64 `json:"private_receipt_account_status,omitempty" xml:"private_receipt_account_status,omitempty"`
	// hotelierIdStatus
	HotelierIdStatus int64 `json:"hotelier_id_status,omitempty" xml:"hotelier_id_status,omitempty"`
	// srlStatus
	SrlStatus int64 `json:"srl_status,omitempty" xml:"srl_status,omitempty"`
	// businessLicenseStatus
	BusinessLicenseStatus int64 `json:"business_license_status,omitempty" xml:"business_license_status,omitempty"`
	// idConfirmLetterStatus
	IdConfirmLetterStatus int64 `json:"id_confirm_letter_status,omitempty" xml:"id_confirm_letter_status,omitempty"`
	// isSigned
	IsSigned int64 `json:"is_signed,omitempty" xml:"is_signed,omitempty"`
	// isAccountApply
	IsAccountApply int64 `json:"is_account_apply,omitempty" xml:"is_account_apply,omitempty"`
	// hid
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// isKezhan
	IsKezhan int64 `json:"is_kezhan,omitempty" xml:"is_kezhan,omitempty"`
	// hotelPicturesStatus
	HotelPicturesStatus int64 `json:"hotel_pictures_status,omitempty" xml:"hotel_pictures_status,omitempty"`
	// canSellStatus
	CanSellStatus int64 `json:"can_sell_status,omitempty" xml:"can_sell_status,omitempty"`
}

var poolDsNhotelInfoDo = sync.Pool{
	New: func() any {
		return new(DsNhotelInfoDo)
	},
}

// GetDsNhotelInfoDo() 从对象池中获取DsNhotelInfoDo
func GetDsNhotelInfoDo() *DsNhotelInfoDo {
	return poolDsNhotelInfoDo.Get().(*DsNhotelInfoDo)
}

// ReleaseDsNhotelInfoDo 释放DsNhotelInfoDo
func ReleaseDsNhotelInfoDo(v *DsNhotelInfoDo) {
	v.OnlineDate = ""
	v.SignDate = ""
	v.Name = ""
	v.OuterId = ""
	v.CanSellDate = ""
	v.OnlineStatus = 0
	v.QualificationType = 0
	v.AccountType = 0
	v.TrainStatus = 0
	v.CommLetterStatus = 0
	v.PrivateReceiptAccountStatus = 0
	v.HotelierIdStatus = 0
	v.SrlStatus = 0
	v.BusinessLicenseStatus = 0
	v.IdConfirmLetterStatus = 0
	v.IsSigned = 0
	v.IsAccountApply = 0
	v.Hid = 0
	v.IsKezhan = 0
	v.HotelPicturesStatus = 0
	v.CanSellStatus = 0
	poolDsNhotelInfoDo.Put(v)
}
