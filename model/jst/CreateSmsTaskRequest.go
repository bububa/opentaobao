package jst

// CreateSmsTaskRequest 
/* model for simplify = false
type CreateSmsTaskRequest struct {

    // 任务Code，系统分配的，创建任务不需要传，更新任务需要传对应的Code（只能更新短信文案、短信签名、短信模板）
    
    TaskCode   string `json:"task_code,omitempty"`
    

    // 任务对应的短信类型 ：1--数字短信  2--权益短信  3--公众号短信
    
    SmsType   int64 `json:"sms_type,omitempty"`
    

    // 权益短信必须是两个文案，其他类型短信为一个文案，文案中必须带${url}占位符，普通短信文案在第一个，权益短信文案在第二个，请严格按照顺序提交
    
    Contents  struct {
        String  []string `json:"string,omitempty"`
    } `json:"contents,omitempty"`
    

    // 数字短信必须是两个模板，其他类型短信为一个模板，普通短信模板在第一个，数字短信的模板在第二个，请严格按照顺序提交
    
    TemplateCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"template_codes,omitempty"`
    

    // 商品详情页或店铺页H5长链地址，不能是短链
    
    Url   string `json:"url,omitempty"`
    

    // 公众号短信必须是两个签名，其他类型短信为一个签名，普通短信的签名在第一个，公众号短信的签名在第二个，请严格按照顺序提交
    
    SignNames  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sign_names,omitempty"`
    

}
*/

// CreateSmsTaskRequest 
type CreateSmsTaskRequest struct {

    // 任务Code，系统分配的，创建任务不需要传，更新任务需要传对应的Code（只能更新短信文案、短信签名、短信模板）
    TaskCode   string `json:"task_code,omitempty"`

    // 任务对应的短信类型 ：1--数字短信  2--权益短信  3--公众号短信
    SmsType   int64 `json:"sms_type,omitempty"`

    // 权益短信必须是两个文案，其他类型短信为一个文案，文案中必须带${url}占位符，普通短信文案在第一个，权益短信文案在第二个，请严格按照顺序提交
    Contents   []string `json:"contents,omitempty"`

    // 数字短信必须是两个模板，其他类型短信为一个模板，普通短信模板在第一个，数字短信的模板在第二个，请严格按照顺序提交
    TemplateCodes   []string `json:"template_codes,omitempty"`

    // 商品详情页或店铺页H5长链地址，不能是短链
    Url   string `json:"url,omitempty"`

    // 公众号短信必须是两个签名，其他类型短信为一个签名，普通短信的签名在第一个，公众号短信的签名在第二个，请严格按照顺序提交
    SignNames   []string `json:"sign_names,omitempty"`

}
