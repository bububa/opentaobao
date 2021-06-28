package jst

// SmsTaskModel 
/* model for simplify = false
type SmsTaskModel struct {

    // appkey
    
    AppKey   string `json:"app_key,omitempty"`
    

    // 短信文案
    
    Contents  struct {
        String  []string `json:"string,omitempty"`
    } `json:"contents,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // isv淘宝nick
    
    IsvNick   string `json:"isv_nick,omitempty"`
    

    // 商家淘宝nick
    
    SellerNick   string `json:"seller_nick,omitempty"`
    

    // 短信签名
    
    SignNames  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sign_names,omitempty"`
    

    // 任务对应的短信类型 ：1--数字短信  2--权益短信  3--公众号短信
    
    SmsType   string `json:"sms_type,omitempty"`
    

    // 系统分配的任务code
    
    TaskCode   string `json:"task_code,omitempty"`
    

    // 短信模板code
    
    TemplateCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"template_codes,omitempty"`
    

    // 商品或店铺详情页H5长链地址
    
    Url   string `json:"url,omitempty"`
    

}
*/

// SmsTaskModel 
type SmsTaskModel struct {

    // appkey
    AppKey   string `json:"app_key,omitempty"`

    // 短信文案
    Contents   []string `json:"contents,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // isv淘宝nick
    IsvNick   string `json:"isv_nick,omitempty"`

    // 商家淘宝nick
    SellerNick   string `json:"seller_nick,omitempty"`

    // 短信签名
    SignNames   []string `json:"sign_names,omitempty"`

    // 任务对应的短信类型 ：1--数字短信  2--权益短信  3--公众号短信
    SmsType   string `json:"sms_type,omitempty"`

    // 系统分配的任务code
    TaskCode   string `json:"task_code,omitempty"`

    // 短信模板code
    TemplateCodes   []string `json:"template_codes,omitempty"`

    // 商品或店铺详情页H5长链地址
    Url   string `json:"url,omitempty"`

}
