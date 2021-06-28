package user

// OpenAccount 
/* model for simplify = false
type OpenAccount struct {

    // 登录名
    
    LoginId   string `json:"login_id,omitempty"`
    

    // 帐号创建的设备的ID
    
    CreateDeviceId   string `json:"create_device_id,omitempty"`
    

    // 支付宝的帐号标识
    
    AlipayId   string `json:"alipay_id,omitempty"`
    

    // 地区
    
    Locale   string `json:"locale,omitempty"`
    

    // 银行卡号
    
    BankCardNo   string `json:"bank_card_no,omitempty"`
    

    // 开发者自定义账号id
    
    IsvAccountId   string `json:"isv_account_id,omitempty"`
    

    // 邮箱
    
    Email   string `json:"email,omitempty"`
    

    // 头像url
    
    AvatarUrl   string `json:"avatar_url,omitempty"`
    

    // 银行卡的拥有者姓名
    
    BankCardOwnerName   string `json:"bank_card_owner_name,omitempty"`
    

    // 展示名
    
    DisplayName   string `json:"display_name,omitempty"`
    

    // 密码salt
    
    LoginPwdSalt   string `json:"login_pwd_salt,omitempty"`
    

    // 密码
    
    LoginPwd   string `json:"login_pwd,omitempty"`
    

    // 第三方oauth openid
    
    OpenId   string `json:"open_id,omitempty"`
    

    // 手机
    
    Mobile   string `json:"mobile,omitempty"`
    

    // 账号创建时的位置
    
    CreateLocation   string `json:"create_location,omitempty"`
    

    // 自定义扩展信息Map的Json格式
    
    ExtInfos   string `json:"ext_infos,omitempty"`
    

    // 密码加密强度
    
    LoginPwdIntensity   int64 `json:"login_pwd_intensity,omitempty"`
    

    // 账号创建类型：1、通过短信创建，2、ISV批量导入，3、ISV OAuth创建
    
    Type   int64 `json:"type,omitempty"`
    

    // 账号状态：1、启用，2、删除，3、禁用
    
    Status   int64 `json:"status,omitempty"`
    

    // 加密算法类型：1、代表单纯MD5，2：代表单一Salt的MD5，3、代表根据记录不同后的MD5
    
    LoginPwdEncryption   int64 `json:"login_pwd_encryption,omitempty"`
    

    // 1男 2女
    
    Gender   int64 `json:"gender,omitempty"`
    

    // 姓名
    
    Name   string `json:"name,omitempty"`
    

    // 出生日期
    
    Birthday   string `json:"birthday,omitempty"`
    

    // 旺旺
    
    Wangwang   string `json:"wangwang,omitempty"`
    

    // 微信
    
    Weixin   string `json:"weixin,omitempty"`
    

    // TAOBAO = 1;WEIXIN = 2;WEIBO = 3;QQ = 4;
    
    OauthPlateform   int64 `json:"oauth_plateform,omitempty"`
    

    // Open Account Id
    
    Id   int64 `json:"id,omitempty"`
    

    // 记录创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 记录上次更新时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 记录的版本号
    
    Version   int64 `json:"version,omitempty"`
    

}
*/

// OpenAccount 
type OpenAccount struct {

    // 登录名
    LoginId   string `json:"login_id,omitempty"`

    // 帐号创建的设备的ID
    CreateDeviceId   string `json:"create_device_id,omitempty"`

    // 支付宝的帐号标识
    AlipayId   string `json:"alipay_id,omitempty"`

    // 地区
    Locale   string `json:"locale,omitempty"`

    // 银行卡号
    BankCardNo   string `json:"bank_card_no,omitempty"`

    // 开发者自定义账号id
    IsvAccountId   string `json:"isv_account_id,omitempty"`

    // 邮箱
    Email   string `json:"email,omitempty"`

    // 头像url
    AvatarUrl   string `json:"avatar_url,omitempty"`

    // 银行卡的拥有者姓名
    BankCardOwnerName   string `json:"bank_card_owner_name,omitempty"`

    // 展示名
    DisplayName   string `json:"display_name,omitempty"`

    // 密码salt
    LoginPwdSalt   string `json:"login_pwd_salt,omitempty"`

    // 密码
    LoginPwd   string `json:"login_pwd,omitempty"`

    // 第三方oauth openid
    OpenId   string `json:"open_id,omitempty"`

    // 手机
    Mobile   string `json:"mobile,omitempty"`

    // 账号创建时的位置
    CreateLocation   string `json:"create_location,omitempty"`

    // 自定义扩展信息Map的Json格式
    ExtInfos   string `json:"ext_infos,omitempty"`

    // 密码加密强度
    LoginPwdIntensity   int64 `json:"login_pwd_intensity,omitempty"`

    // 账号创建类型：1、通过短信创建，2、ISV批量导入，3、ISV OAuth创建
    Type   int64 `json:"type,omitempty"`

    // 账号状态：1、启用，2、删除，3、禁用
    Status   int64 `json:"status,omitempty"`

    // 加密算法类型：1、代表单纯MD5，2：代表单一Salt的MD5，3、代表根据记录不同后的MD5
    LoginPwdEncryption   int64 `json:"login_pwd_encryption,omitempty"`

    // 1男 2女
    Gender   int64 `json:"gender,omitempty"`

    // 姓名
    Name   string `json:"name,omitempty"`

    // 出生日期
    Birthday   string `json:"birthday,omitempty"`

    // 旺旺
    Wangwang   string `json:"wangwang,omitempty"`

    // 微信
    Weixin   string `json:"weixin,omitempty"`

    // TAOBAO = 1;WEIXIN = 2;WEIBO = 3;QQ = 4;
    OauthPlateform   int64 `json:"oauth_plateform,omitempty"`

    // Open Account Id
    Id   int64 `json:"id,omitempty"`

    // 记录创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 记录上次更新时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 记录的版本号
    Version   int64 `json:"version,omitempty"`

}
