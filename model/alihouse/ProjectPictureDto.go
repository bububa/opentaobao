package alihouse

// ProjectPictureDto 
type ProjectPictureDto struct {
    // 是否删除 1 是 0 否
    IsDeleted   string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
    // 说明描述
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 图片全路径
    PicData   string `json:"pic_data,omitempty" xml:"pic_data,omitempty"`
    // 图片分类
    Category   int64 `json:"category,omitempty" xml:"category,omitempty"`
    // 图片排序值
    OrderTag   int64 `json:"order_tag,omitempty" xml:"order_tag,omitempty"`
    // 外部户型ID
    OuterLayoutId   string `json:"outer_layout_id,omitempty" xml:"outer_layout_id,omitempty"`
    // 楼盘ID
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 图片唯一ID
    OuterPictureId   string `json:"outer_picture_id,omitempty" xml:"outer_picture_id,omitempty"`
    // 图片名称
    PicName   string `json:"pic_name,omitempty" xml:"pic_name,omitempty"`
}
