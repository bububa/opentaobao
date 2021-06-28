package taotv

// TaobaoTaotvCarouselCategoryListModel 
/* model for simplify = false
type TaobaoTaotvCarouselCategoryListModel struct {

    // 分类排序
    
    Sort   int64 `json:"sort,omitempty"`
    

    // 分类图片
    
    Pic   string `json:"pic,omitempty"`
    

    // 分类牌照方
    
    Bcp   int64 `json:"bcp,omitempty"`
    

    // 分类名称
    
    Name   string `json:"name,omitempty"`
    

    // 分类ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 分类频道列表
    
    ChannelList  struct {
        Channels  []Channels `json:"channels,omitempty"`
    } `json:"channel_list,omitempty"`
    

}
*/

// TaobaoTaotvCarouselCategoryListModel 
type TaobaoTaotvCarouselCategoryListModel struct {

    // 分类排序
    Sort   int64 `json:"sort,omitempty"`

    // 分类图片
    Pic   string `json:"pic,omitempty"`

    // 分类牌照方
    Bcp   int64 `json:"bcp,omitempty"`

    // 分类名称
    Name   string `json:"name,omitempty"`

    // 分类ID
    Id   int64 `json:"id,omitempty"`

    // 分类频道列表
    ChannelList   []Channels `json:"channel_list,omitempty"`

}
