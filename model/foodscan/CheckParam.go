package foodscan

// CheckParam 
type CheckParam struct {
    // 图片base64编码，注意字符串前缀
    BaseCode   string `json:"base_code,omitempty" xml:"base_code,omitempty"`
    // 1男2女
    Gender   int64 `json:"gender,omitempty" xml:"gender,omitempty"`
    // 淘宝的nickName
    NickName   string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
    // 前8位是日期，后10位是随机字符串
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 第几次拍照 0 左脚背面 1 左脚内侧 2 左脚外侧 3右脚脚背 4右脚内侧 5右脚外侧
    Index   int64 `json:"index,omitempty" xml:"index,omitempty"`
    // 1左脚 2右脚
    ModelType   int64 `json:"model_type,omitempty" xml:"model_type,omitempty"`
    // 用户唯一标识，可以是淘宝用户Id
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
