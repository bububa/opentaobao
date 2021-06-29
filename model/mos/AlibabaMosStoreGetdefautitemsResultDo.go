package mos

// AlibabaMosStoreGetdefautitemsResultDO 
type AlibabaMosStoreGetdefautitemsResultDO struct {
    // data
    Data   *MjStoreItemsTopVo `json:"data,omitempty" xml:"data,omitempty"`
    // errCode
    ErrCode   int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
    // errMsg
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // total
    Total   int64 `json:"total,omitempty" xml:"total,omitempty"`
}
