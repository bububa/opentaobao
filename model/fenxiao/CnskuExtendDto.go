package fenxiao

// CnskuExtendDto 
type CnskuExtendDto struct {
    // 配送要求（1：顺丰优先配）
    DeliverRequirements   string `json:"deliver_requirements,omitempty" xml:"deliver_requirements,omitempty"`
    // 是否贵品
    IsPrecious   bool `json:"is_precious,omitempty" xml:"is_precious,omitempty"`
    // 商品图片
    PictureUrl   string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
    // sn 示例
    CnskuSnSampleDTOList   []Null `json:"cnsku_sn_sample_d_t_o_list,omitempty" xml:"cnsku_sn_sample_d_t_o_list>null,omitempty"`
    // 主要成分
    MainComposition   string `json:"main_composition,omitempty" xml:"main_composition,omitempty"`
    // 是否进口
    IsImported   bool `json:"is_imported,omitempty" xml:"is_imported,omitempty"`
    // 运输单元宽
    TransWidth   int64 `json:"trans_width,omitempty" xml:"trans_width,omitempty"`
    // 包装单位
    PackageUnit   string `json:"package_unit,omitempty" xml:"package_unit,omitempty"`
    // 温度标识 1:常温 2:5°C-12°C 3:0°C-4°C 4:-18°C-0°C
    TemperatureRequirement   string `json:"temperature_requirement,omitempty" xml:"temperature_requirement,omitempty"`
    // 运输单元长
    TransLength   int64 `json:"trans_length,omitempty" xml:"trans_length,omitempty"`
    // 是否包含电池(默认无电池) 1:无电池, 2:内置电池, 3:外置电池
    IncludeBattery   string `json:"include_battery,omitempty" xml:"include_battery,omitempty"`
    // 商品使用的包材材质
    PackageMaterialClass   string `json:"package_material_class,omitempty" xml:"package_material_class,omitempty"`
    // 3w货品分类属性（大电用）空调属性，内外机
    TypeProperty   string `json:"type_property,omitempty" xml:"type_property,omitempty"`
    // 生产厂家
    Manufacturer   string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
    // 内长
    InnerLength   int64 `json:"inner_length,omitempty" xml:"inner_length,omitempty"`
    // 运输单元体积照片
    TransImageUrl   string `json:"trans_image_url,omitempty" xml:"trans_image_url,omitempty"`
    // 承重
    LoadBearing   int64 `json:"load_bearing,omitempty" xml:"load_bearing,omitempty"`
    // 货品图片
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    // 是否生产批号管理
    IsProduceCodeMgt   bool `json:"is_produce_code_mgt,omitempty" xml:"is_produce_code_mgt,omitempty"`
    // 是否监管
    IsDrugs   bool `json:"is_drugs,omitempty" xml:"is_drugs,omitempty"`
    // 原产国
    BrandCountry   string `json:"brand_country,omitempty" xml:"brand_country,omitempty"`
    // 币种
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    // 产地
    ProducingArea   string `json:"producing_area,omitempty" xml:"producing_area,omitempty"`
    // 运输单元高
    TransHeight   int64 `json:"trans_height,omitempty" xml:"trans_height,omitempty"`
    // 运输单元体积
    TransVolume   int64 `json:"trans_volume,omitempty" xml:"trans_volume,omitempty"`
    // 内高
    InnerHeight   int64 `json:"inner_height,omitempty" xml:"inner_height,omitempty"`
    // 备案链接
    RecordUrl   string `json:"record_url,omitempty" xml:"record_url,omitempty"`
    // 运输单元重量
    TransWeight   int64 `json:"trans_weight,omitempty" xml:"trans_weight,omitempty"`
    // 包材类型 1:空白箱 2：菜鸟联盟包材
    MaterialType   int64 `json:"material_type,omitempty" xml:"material_type,omitempty"`
    // 商品使用的包材范围 1：商家包材 2：CP包材
    PackageMaterialMode   int64 `json:"package_material_mode,omitempty" xml:"package_material_mode,omitempty"`
    // 存储分类
    Classification   string `json:"classification,omitempty" xml:"classification,omitempty"`
    // 税收分类编码
    TaxCode   string `json:"tax_code,omitempty" xml:"tax_code,omitempty"`
    // 大电barCode
    WwwBarCode   string `json:"www_bar_code,omitempty" xml:"www_bar_code,omitempty"`
    // 税率标示
    TaxRate   string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
    // 内宽
    InnerWidth   int64 `json:"inner_width,omitempty" xml:"inner_width,omitempty"`
    // 剂型
    DosageForms   string `json:"dosage_forms,omitempty" xml:"dosage_forms,omitempty"`
}
