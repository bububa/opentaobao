package shop

// Shop 
type Shop struct {

    // 店铺编号
    Sid   int64 `json:"sid,omitempty"`

    // 店铺标题
    Title   string `json:"title,omitempty"`

    // 店标地址。返回相对路径，可以用"http://logo.taobao.com/shop-logo"来拼接成绝对路径
    PicPath   string `json:"pic_path,omitempty"`

    // 最后修改时间。格式：yyyy-MM-dd HH:mm:ss
    Modified   string `json:"modified,omitempty"`

}
