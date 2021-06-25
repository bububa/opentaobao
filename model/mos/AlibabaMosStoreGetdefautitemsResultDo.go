package mos

// AlibabaMosStoreGetdefautitemsResultDo 
type AlibabaMosStoreGetdefautitemsResultDo struct {

    // data
    Data   *MjStoreItemsTopVo `json:"data,omitempty"`

    // errCode
    ErrCode   int64 `json:"err_code,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // total
    Total   int64 `json:"total,omitempty"`

}
