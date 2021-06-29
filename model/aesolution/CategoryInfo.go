package aesolution

// CategoryInfo 
type CategoryInfo struct {
    // category id
    ChildrenCategoryId   int64 `json:"children_category_id,omitempty" xml:"children_category_id,omitempty"`
    // whether the category is leaf or not
    IsLeafCategory   bool `json:"is_leaf_category,omitempty" xml:"is_leaf_category,omitempty"`
    // level of the categories. As for root categories, the level is 1
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    // multi langauge names of the categories
    MultiLanguageNames   string `json:"multi_language_names,omitempty" xml:"multi_language_names,omitempty"`
}
