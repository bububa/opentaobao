package jst

// JdpUser 
type JdpUser struct {

    // 用户昵称
    UserNick   string `json:"user_nick,omitempty"`

    // rds数据库的实例名
    RdsName   string `json:"rds_name,omitempty"`

    // 0:暂停1：正常2：sessoin失效，停止3：已删除
    Status   int64 `json:"status,omitempty"`

}
