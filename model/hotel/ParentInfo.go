package hotel

// ParentInfo 
/* model for simplify = false
type ParentInfo struct {

    // 脱敏后的userId
    
    UserId   int64 `json:"user_id,omitempty"`
    

    // 脱敏后的用户名字
    
    UserNick   string `json:"user_nick,omitempty"`
    

}
*/

// ParentInfo 
type ParentInfo struct {

    // 脱敏后的userId
    UserId   int64 `json:"user_id,omitempty"`

    // 脱敏后的用户名字
    UserNick   string `json:"user_nick,omitempty"`

}
