package xiamicontent

// SongCatTagDto 
type SongCatTagDto struct {
    // tag间的关系查询 1 and 2：or
    Relation   int64 `json:"relation,omitempty" xml:"relation,omitempty"`
    // tag code列表
    TagCodes   []string `json:"tag_codes,omitempty" xml:"tag_codes>string,omitempty"`
}
