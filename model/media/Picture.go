package media

// Picture 
type Picture struct {

    // 图片ID
    PictureId   int64 `json:"picture_id,omitempty"`

    // 图片分类ID
    PictureCategoryId   int64 `json:"picture_category_id,omitempty"`

    // 返回的是绝对路径如：http://img07.taobaocdn.com/imgextra/i7/22670458/T2dD0kXb4cXXXXXXXX_!!22670458.jpg
    PicturePath   string `json:"picture_path,omitempty"`

    // 图片标题
    Title   string `json:"title,omitempty"`

    // 图片大小,bite单位
    Sizes   int64 `json:"sizes,omitempty"`

    // 图片相素,格式:长x宽，如450x150
    Pixel   string `json:"pixel,omitempty"`

    // 图片状态,unfroze代表没有被冻结，froze代表被冻结,pass代表排查通过
    Status   string `json:"status,omitempty"`

    // 图片是否删除的标记
    Deleted   string `json:"deleted,omitempty"`

    // 图片的创建时间
    Created   string `json:"created,omitempty"`

    // 图片的修改时间
    Modified   string `json:"modified,omitempty"`

    // 图片是否被引用
    Referenced   bool `json:"referenced,omitempty"`

    // 图片在后台处理之后的md5值当md5值为32位长度的字符串时为图片搬家后的文件md5验证码md5值为长整数时为图片替换后的时间戳
    Md5   string `json:"md5,omitempty"`

    // 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布
    ClientType   string `json:"client_type,omitempty"`

}
