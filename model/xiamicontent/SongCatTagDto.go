package xiamicontent

// SongCatTagDto 结构体
type SongCatTagDto struct {
	// tag code列表
	TagCodes []string `json:"tag_codes,omitempty" xml:"tag_codes>string,omitempty"`
	// tag间的关系查询 1 and 2：or
	Relation int64 `json:"relation,omitempty" xml:"relation,omitempty"`
}
