package tmallcarenter

// OriginVehicleInfoDto 
type OriginVehicleInfoDto struct {
    // 销售名称
    SalesName   string `json:"sales_name,omitempty" xml:"sales_name,omitempty"`
    // 生产年份
    ProductiveYear   int64 `json:"productive_year,omitempty" xml:"productive_year,omitempty"`
    // 厂家类型(国产,合资,进口)
    ManufactureType   string `json:"manufacture_type,omitempty" xml:"manufacture_type,omitempty"`
    // 后轮胎规格
    RearTyre   string `json:"rear_tyre,omitempty" xml:"rear_tyre,omitempty"`
    // 品牌名称
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    // 底盘压缩ID
    ChassisCid   string `json:"chassis_cid,omitempty" xml:"chassis_cid,omitempty"`
    // 上市年份
    ReleaseYear   int64 `json:"release_year,omitempty" xml:"release_year,omitempty"`
    // 厂家名称
    ManufactureName   string `json:"manufacture_name,omitempty" xml:"manufacture_name,omitempty"`
    // 销售年款
    SalesYear   int64 `json:"sales_year,omitempty" xml:"sales_year,omitempty"`
    // 前轮胎规格
    FrontTyre   string `json:"front_tyre,omitempty" xml:"front_tyre,omitempty"`
    // 排量
    Displacement   string `json:"displacement,omitempty" xml:"displacement,omitempty"`
    // 销售状态
    SalesStatus   string `json:"sales_status,omitempty" xml:"sales_status,omitempty"`
    // 车系
    LineName   string `json:"line_name,omitempty" xml:"line_name,omitempty"`
    // 排放标准
    Emmission   string `json:"emmission,omitempty" xml:"emmission,omitempty"`
    // 生产状态
    ProductiveStatus   string `json:"productive_status,omitempty" xml:"productive_status,omitempty"`
    // 停产年份
    EndYear   int64 `json:"end_year,omitempty" xml:"end_year,omitempty"`
    // 状态(状态,0:删除,1:新增,2:修改)
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 发动起型号
    EngineModel   string `json:"engine_model,omitempty" xml:"engine_model,omitempty"`
    // 扩展信息(json串)
    ExtendInfo   string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
    // 最大功率(KW)
    MaxPower   string `json:"max_power,omitempty" xml:"max_power,omitempty"`
    // 国别
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    // 销售版本
    SalesEdition   string `json:"sales_edition,omitempty" xml:"sales_edition,omitempty"`
    // 版本压缩id
    VersionCid   string `json:"version_cid,omitempty" xml:"version_cid,omitempty"`
    // 车型id(主键字段,代表唯一一条数据)
    VehicleInfoId   string `json:"vehicle_info_id,omitempty" xml:"vehicle_info_id,omitempty"`
    // 车辆级别
    VehicleLevel   string `json:"vehicle_level,omitempty" xml:"vehicle_level,omitempty"`
    // 官方指导价(万元)
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // 车型
    ModelName   string `json:"model_name,omitempty" xml:"model_name,omitempty"`
    // 上市月份
    ReleaseMonth   int64 `json:"release_month,omitempty" xml:"release_month,omitempty"`
    // 燃油类型
    FuelType   string `json:"fuel_type,omitempty" xml:"fuel_type,omitempty"`
    // 变速箱类型
    TransmissionType   string `json:"transmission_type,omitempty" xml:"transmission_type,omitempty"`
    // 驱动方式
    DriveModel   string `json:"drive_model,omitempty" xml:"drive_model,omitempty"`
    // 代号/底盘型号
    ChassisNum   string `json:"chassis_num,omitempty" xml:"chassis_num,omitempty"`
    // 车身类型
    BodyModel   string `json:"body_model,omitempty" xml:"body_model,omitempty"`
    // 车辆类型
    VehicleType   string `json:"vehicle_type,omitempty" xml:"vehicle_type,omitempty"`
    // 代数
    GenerationNum   string `json:"generation_num,omitempty" xml:"generation_num,omitempty"`
}
