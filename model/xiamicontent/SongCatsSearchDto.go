package xiamicontent

// SongCatsSearchDto 
type SongCatsSearchDto struct {

    // 标签类目间的关系查询 1 and 2：or
    
    Relation   int64 `json:"relation,omitempty" xml:"relation,omitempty"`
    

    // 标签组列表
    
    CatTagCodes   []SongCatTagDto `json:"cat_tag_codes,omitempty" xml:"cat_tag_codes,omitempty"`
    

}
