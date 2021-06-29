package moscm

// SpuInputDTO 
type SpuInputDTO struct {
    // 产品条码信息
    BarcodeStr   string `json:"barcode_str,omitempty" xml:"barcode_str,omitempty"`
    // 银泰品牌Id
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // 品牌名称
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    // 商品类目名称
    CatName   string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
    // 商品类目ID.必须是叶子类目ID
    Cid   string `json:"cid,omitempty" xml:"cid,omitempty"`
    // SPU ID
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // 是否新品默认是true
    IsNew   bool `json:"is_new,omitempty" xml:"is_new,omitempty"`
    // 透明素材图
    Material   string `json:"material,omitempty" xml:"material,omitempty"`
    // 已废弃
    Mdesc   string `json:"mdesc,omitempty" xml:"mdesc,omitempty"`
    // 已废弃
    PcDesc   string `json:"pc_desc,omitempty" xml:"pc_desc,omitempty"`
    // 产品的主图片地址.(绝对地址,格式:http://host/image_path)
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    // 吊牌价.单位为元.
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // 产品ID
    ProductId   string `json:"product_id,omitempty" xml:"product_id,omitempty"`
    // 产品的子图片.目前最多支持50张。
    ProductImgs   []ProductImgDTO `json:"product_imgs,omitempty" xml:"product_imgs>product_img_dto,omitempty"`
    // 产品参数描述，如：材质成分、裙长、年份季节等信息，最多50个属性
    Props   []PropertyDTO `json:"props,omitempty" xml:"props>property_dto,omitempty"`
    // 产品卖点描述，长度限制20个汉字
    SellPt   string `json:"sell_pt,omitempty" xml:"sell_pt,omitempty"`
    // 款号
    StyleNo   string `json:"style_no,omitempty" xml:"style_no,omitempty"`
    // 子标题
    SubTitle   string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
    // brand是银泰品牌Id，colorName（颜色名称）、colorCode(颜色编码)、sizeCode(尺码编码)、sizeName(尺码名称):商品销售属性，标签一些在属性里找不到id的属性存放在这里
    Tags   string `json:"tags,omitempty" xml:"tags,omitempty"`
    // 产品名称
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 天猫品牌Id
    TmallBrandId   string `json:"tmall_brand_id,omitempty" xml:"tmall_brand_id,omitempty"`
    // 详情图（最大列表长度：60）
    DescPicList   []string `json:"desc_pic_list,omitempty" xml:"desc_pic_list>string,omitempty"`
    // 天猫ItemId
    TmallItemId   int64 `json:"tmall_item_id,omitempty" xml:"tmall_item_id,omitempty"`
}
