package baodian

// UserCoinDeposit 
type UserCoinDeposit struct {

    // 用户AuthCode
    UserAuthCode   string `json:"userAuthCode,omitempty"`

    // 宝点用户体系中的用户字符串ID
    UserStrId   string `json:"user_str_id,omitempty"`

    // 在数娱用户体系中的用户昵称，如果未设置则默认返回淘宝的用户名
    UserNick   string `json:"user_nick,omitempty"`

    // 当前宝点帐户余额，即当前宝点数
    Deposit   int64 `json:"deposit,omitempty"`

    // 当前信用赊欠宝点数
    Credit   int64 `json:"credit,omitempty"`

    // 当前的信用帐期，即还款期截止日
    CreditPeriod   string `json:"credit_period,omitempty"`

    // 用户是否有逾期宝点贷款
    IsExpired   bool `json:"is_expired,omitempty"`

    // 用户当前可用的信用消费的额度，即还可以赊欠多少宝点
    CreditLimit   int64 `json:"credit_limit,omitempty"`

    // 单位数量(10个宝点)宝点价格，即10个宝点的价格
    Price   int64 `json:"price,omitempty"`

    // 能否支付
    EnablePay   bool `json:"enable_pay,omitempty"`

    // 是否新注册用户
    NewUser   bool `json:"new_user,omitempty"`

    // 用户积分数量
    Gamepoints   int64 `json:"gamepoints,omitempty"`

}
