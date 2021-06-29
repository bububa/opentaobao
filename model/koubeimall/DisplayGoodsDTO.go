package koubeimall

// DisplayGoodsDTO 
type DisplayGoodsDTO struct {
    // 是否有图模式。true：有图，false：无图
    HasPicture   bool `json:"has_picture,omitempty" xml:"has_picture,omitempty"`
    // 推荐菜数量
    GoodsCount   int64 `json:"goods_count,omitempty" xml:"goods_count,omitempty"`
    // 带图片的推荐菜详情模型
    GoodsDetailInfos   []GoodsDetailInfo `json:"goods_detail_infos,omitempty" xml:"goods_detail_infos>goods_detail_info,omitempty"`
    // 标题：推荐菜
    GoodsTitle   string `json:"goods_title,omitempty" xml:"goods_title,omitempty"`
    // 无图模式，菜品文案
    GoodsDesc   string `json:"goods_desc,omitempty" xml:"goods_desc,omitempty"`
}
