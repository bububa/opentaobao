package tmc

// TmcUser 
/* model for simplify = false
type TmcUser struct {

    // 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户
    
    UserPlatform   string `json:"user_platform,omitempty"`
    

    // 用户昵称
    
    UserNick   string `json:"user_nick,omitempty"`
    

    // 用户ID
    
    UserId   int64 `json:"user_id,omitempty"`
    

    // 用户授权是否有效，true表示授权有效，false表示授权过期
    
    IsValid   bool `json:"is_valid,omitempty"`
    

    // 用户首次开通时间
    
    Created   string `json:"created,omitempty"`
    

    // 用户最后开通时间
    
    Modified   string `json:"modified,omitempty"`
    

    // 用户开通的消息类型列表。如果为空表示应用开通的所有类型
    
    Topics  struct {
        String  []string `json:"string,omitempty"`
    } `json:"topics,omitempty"`
    

    // 接收用户消息的组名
    
    GroupName   string `json:"group_name,omitempty"`
    

}
*/

// TmcUser 
type TmcUser struct {

    // 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户
    UserPlatform   string `json:"user_platform,omitempty"`

    // 用户昵称
    UserNick   string `json:"user_nick,omitempty"`

    // 用户ID
    UserId   int64 `json:"user_id,omitempty"`

    // 用户授权是否有效，true表示授权有效，false表示授权过期
    IsValid   bool `json:"is_valid,omitempty"`

    // 用户首次开通时间
    Created   string `json:"created,omitempty"`

    // 用户最后开通时间
    Modified   string `json:"modified,omitempty"`

    // 用户开通的消息类型列表。如果为空表示应用开通的所有类型
    Topics   []string `json:"topics,omitempty"`

    // 接收用户消息的组名
    GroupName   string `json:"group_name,omitempty"`

}
