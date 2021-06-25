package taotv

// TaobaoTaotvVideoPlaylistPageModel 
type TaobaoTaotvVideoPlaylistPageModel struct {

    // 当前页
    PageNo   int64 `json:"page_no,omitempty"`

    // 此接口默认每次获取100条
    PageSize   int64 `json:"page_size,omitempty"`

    // 播单总数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 总共页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 播单对象
    DataList   []TaobaoTaotvVideoPlaylistPageData `json:"data_list,omitempty"`

}
