package xhotel

// DchotelSignDo 结构体
type DchotelSignDo struct {
	// taccAuthLetterAllowDate
	TaccAuthLetterAllowDate string `json:"tacc_auth_letter_allow_date,omitempty" xml:"tacc_auth_letter_allow_date,omitempty"`
	// phone
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// email
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// intentAchieveDate
	IntentAchieveDate string `json:"intent_achieve_date,omitempty" xml:"intent_achieve_date,omitempty"`
	// hotelAccount
	HotelAccount string `json:"hotel_account,omitempty" xml:"hotel_account,omitempty"`
	// privateReceiptAccountStatus
	PrivateReceiptAccountStatus int64 `json:"private_receipt_account_status,omitempty" xml:"private_receipt_account_status,omitempty"`
	// hotelierIdStatus
	HotelierIdStatus int64 `json:"hotelier_id_status,omitempty" xml:"hotelier_id_status,omitempty"`
	// stlStatus
	StlStatus int64 `json:"stl_status,omitempty" xml:"stl_status,omitempty"`
	// blStatus
	BlStatus int64 `json:"bl_status,omitempty" xml:"bl_status,omitempty"`
	// idConfirmLetterStatus
	IdConfirmLetterStatus int64 `json:"id_confirm_letter_status,omitempty" xml:"id_confirm_letter_status,omitempty"`
	// qlfType
	QlfType int64 `json:"qlf_type,omitempty" xml:"qlf_type,omitempty"`
	// payType
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// isKezhanSign
	IsKezhanSign int64 `json:"is_kezhan_sign,omitempty" xml:"is_kezhan_sign,omitempty"`
	// signStatus
	SignStatus int64 `json:"sign_status,omitempty" xml:"sign_status,omitempty"`
	// pid
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// hid
	Hid int64 `json:"hid,omitempty" xml:"hid,omitempty"`
	// hotelPicturesStatus
	HotelPicturesStatus int64 `json:"hotel_pictures_status,omitempty" xml:"hotel_pictures_status,omitempty"`
	// commLetterStatus
	CommLetterStatus int64 `json:"comm_letter_status,omitempty" xml:"comm_letter_status,omitempty"`
	// toDcHotelStatus
	ToDcHotelStatus int64 `json:"to_dc_hotel_status,omitempty" xml:"to_dc_hotel_status,omitempty"`
}
