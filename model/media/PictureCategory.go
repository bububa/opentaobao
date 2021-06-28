package media

// PictureCategory 
/* model for simplify = false
type PictureCategory struct {

    // 图片分类ID
    
    PictureCategoryId   int64 `json:"picture_category_id,omitempty"`
    

    // 图片分类名
    
    PictureCategoryName   string `json:"picture_category_name,omitempty"`
    

    // 图片分类型别，sys-fixture代表店铺装修分类(系统分类)，sys-auction代表宝贝图片分类(系统分类)，user-define代表用户自定义分类
    
    Type   string `json:"type,omitempty"`
    

    // 图片类目的创建时间
    
    Created   string `json:"created,omitempty"`
    

    // 图片分类的修改时间
    
    Modified   string `json:"modified,omitempty"`
    

    // 图片分类排序
    
    Position   int64 `json:"position,omitempty"`
    

    // 一级分类的parent_id为0<br/>二级分类的则为其父分类的picture_category_id
    
    ParentId   int64 `json:"parent_id,omitempty"`
    

}
*/

// PictureCategory 
type PictureCategory struct {

    // 图片分类ID
    PictureCategoryId   int64 `json:"picture_category_id,omitempty"`

    // 图片分类名
    PictureCategoryName   string `json:"picture_category_name,omitempty"`

    // 图片分类型别，sys-fixture代表店铺装修分类(系统分类)，sys-auction代表宝贝图片分类(系统分类)，user-define代表用户自定义分类
    Type   string `json:"type,omitempty"`

    // 图片类目的创建时间
    Created   string `json:"created,omitempty"`

    // 图片分类的修改时间
    Modified   string `json:"modified,omitempty"`

    // 图片分类排序
    Position   int64 `json:"position,omitempty"`

    // 一级分类的parent_id为0<br/>二级分类的则为其父分类的picture_category_id
    ParentId   int64 `json:"parent_id,omitempty"`

}
