package campus

// IdentifyAuthDto 
type IdentifyAuthDto struct {
    // app_code
    AppCode   string `json:"app_code,omitempty" xml:"app_code,omitempty"`
    // device_id
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    // 鉴权类型
    IdentifAuthTypeEnum   string `json:"identif_auth_type_enum,omitempty" xml:"identif_auth_type_enum,omitempty"`
    // app_sigin
    AppSign   string `json:"app_sign,omitempty" xml:"app_sign,omitempty"`
    // 时间
    TimeStamp   int64 `json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
    // 凭证
    VoucherDTOList   []VoucherDto `json:"voucher_d_t_o_list,omitempty" xml:"voucher_d_t_o_list>voucher_dto,omitempty"`
}
