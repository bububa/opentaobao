package xiami

// CollectDetail 
type CollectDetail struct {

    // 是否默认图片
    LogoDefault   bool `json:"logo_default,omitempty"`

    // 精选集ID
    ListId   int64 `json:"list_id,omitempty"`

    // 创建者用户ID
    UserId   int64 `json:"user_id,omitempty"`

    // 精选集名称
    CollectName   string `json:"collect_name,omitempty"`

    // 精选集LOGO
    Logo   string `json:"logo,omitempty"`

    // 精选集介绍
    Description   string `json:"description,omitempty"`

    // 歌曲总数
    SongsCount   int64 `json:"songs_count,omitempty"`

    // 试听数
    PlayCount   int64 `json:"play_count,omitempty"`

    // 创建精选集时间戳
    GmtCreate   int64 `json:"gmt_create,omitempty"`

    // 修改精选集时间戳
    GmtModify   int64 `json:"gmt_modify,omitempty"`

    // 浏览次数
    Views   int64 `json:"views,omitempty"`

    // 下载次数
    Downloads   int64 `json:"downloads,omitempty"`

    // 喜欢次数
    Favorites   int64 `json:"favorites,omitempty"`

    // 评论次数
    Comments   int64 `json:"comments,omitempty"`

    // 推荐次数
    Recommends   int64 `json:"recommends,omitempty"`

    // 创建者用户昵称
    UserName   string `json:"user_name,omitempty"`

    // 创建者用户头像
    AuthorAvatar   string `json:"author_avatar,omitempty"`

    // 创建者头像是否默认
    AvatarDefault   bool `json:"avatar_default,omitempty"`

    // 创建者是否VIP
    IsVip   int64 `json:"is_vip,omitempty"`

    // 精选集歌曲列表
    Songs   []CollectSong `json:"songs,omitempty"`

}
