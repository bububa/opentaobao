package iot

// AiCloudResult 
type AiCloudResult struct {

    // 总共符合数据的个数（目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空）
    RecordCount   int64 `json:"record_count,omitempty"`

    // List数据结构，List中的每一个条目是一个Map对象，key为属性名称，value为对应的值
    Model   string `json:"model,omitempty"`

    // 返回码
    MsgCode   string `json:"msg_code,omitempty"`

    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 请求是否成功
    Success   bool `json:"success,omitempty"`

    // 额外扩展信息，JSON串，不一定有
    ExtendInfo   string `json:"extend_info,omitempty"`

    // auth信息
    AuthInfo   string `json:"auth_info,omitempty"`

    // 是否操作成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // uuid 32位
    Uuids   []String `json:"uuids,omitempty"`

    // 对话流列表数据
    Models   []Json `json:"models,omitempty"`

    // List<Map<String, Object>>数据结构，List中的每一个条目是一个Map对象，key为属性名称，value为对应的值
    ModelList   []Json `json:"model_list,omitempty"`

    // model
    Likes   []Like `json:"likes,omitempty"`

}
