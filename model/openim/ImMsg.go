package openim

// ImMsg 
type ImMsg struct {

    // 消息发送者
    FromUser   string `json:"from_user,omitempty"`

    // 消息接受者
    ToUsers   []String `json:"to_users,omitempty"`

    // 消息类型。0:文本消息。1:图片消息，只支持jpg、gif。2:语音消息，只支持amr。8:地理位置信息。
    MsgType   int64 `json:"msg_type,omitempty"`

    // 发送的消息内容。根据不同消息类型，传不同的值。0(文本消息):填消息内容字符串。1(图片):base64编码的jpg或gif文件。3(语音):base64编码的amr文件。8(地理位置):经纬度，格式如 111,222
    Context   string `json:"context,omitempty"`

    // 接收方appkey，默认本app，跨app发送时需要用到
    ToAppkey   string `json:"to_appkey,omitempty"`

    // json map，媒体信息属性。根据msgtype变化。0(文本):填空即可。 1(图片):需要图片格式，{"type":"jpg"}或{"type":"gif"}。   2(语音): 需要文件格式和语音长度信息{"type":"amr","playtime":5}
    MediaAttr   string `json:"media_attr,omitempty"`

    // 如果为1，则表示发送方是一个淘宝账号，该账号必须是本appkey绑定过的客服账号，并且只能给本appkey的用户发消息。通过该参数可以从服务端发起一个客服到用户的会话。
    FromTaobao   int64 `json:"from_taobao,omitempty"`

}
