package xhotel

// DchotelSignDo 
type DchotelSignDo struct {

    // hid
    Hid   int64 `json:"hid,omitempty"`

    // pid
    Pid   int64 `json:"pid,omitempty"`

    // signStatus
    SignStatus   int64 `json:"sign_status,omitempty"`

    // isKezhanSign
    IsKezhanSign   int64 `json:"is_kezhan_sign,omitempty"`

    // name
    Name   string `json:"name,omitempty"`

    // email
    Email   string `json:"email,omitempty"`

    // phone
    Phone   string `json:"phone,omitempty"`

    // payType
    PayType   int64 `json:"pay_type,omitempty"`

    // qlfType
    QlfType   int64 `json:"qlf_type,omitempty"`

    // idConfirmLetterStatus
    IdConfirmLetterStatus   int64 `json:"id_confirm_letter_status,omitempty"`

    // taccAuthLetterAllowDate
    TaccAuthLetterAllowDate   string `json:"tacc_auth_letter_allow_date,omitempty"`

    // stlStatus
    StlStatus   int64 `json:"stl_status,omitempty"`

    // hotelierIdStatus
    HotelierIdStatus   int64 `json:"hotelier_id_status,omitempty"`

    // commLetterStatus
    CommLetterStatus   int64 `json:"comm_letter_status,omitempty"`

    // hotelPicturesStatus
    HotelPicturesStatus   int64 `json:"hotel_pictures_status,omitempty"`

    // hotelAccount
    HotelAccount   string `json:"hotel_account,omitempty"`

    // toDcHotelStatus
    ToDcHotelStatus   int64 `json:"to_dc_hotel_status,omitempty"`

    // intentAchieveDate
    IntentAchieveDate   string `json:"intent_achieve_date,omitempty"`

    // privateReceiptAccountStatus
    PrivateReceiptAccountStatus   int64 `json:"private_receipt_account_status,omitempty"`

    // blStatus
    BlStatus   int64 `json:"bl_status,omitempty"`

}
