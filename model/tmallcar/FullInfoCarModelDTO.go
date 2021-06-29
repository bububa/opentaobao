package tmallcar

// FullInfoCarModelDTO 
type FullInfoCarModelDTO struct {
    // ABS防抱死
    Abs   string `json:"abs,omitempty" xml:"abs,omitempty"`
    // 空调
    Ac   string `json:"ac,omitempty" xml:"ac,omitempty"`
    // 温度分区控制
    AcZoneControl   string `json:"ac_zone_control,omitempty" xml:"ac_zone_control,omitempty"`
    // 自适应巡航
    Acc   string `json:"acc,omitempty" xml:"acc,omitempty"`
    // 加速时间（0-100Km/h）
    Acceleration   string `json:"acceleration,omitempty" xml:"acceleration,omitempty"`
    // 离去角
    AccessAngle   string `json:"access_angle,omitempty" xml:"access_angle,omitempty"`
    // 主动刹车
    ActiveBrake   string `json:"active_brake,omitempty" xml:"active_brake,omitempty"`
    // 主动转向系统
    ActiveSteering   string `json:"active_steering,omitempty" xml:"active_steering,omitempty"`
    // 自动头灯
    AdaptiveHeadlamp   string `json:"adaptive_headlamp,omitempty" xml:"adaptive_headlamp,omitempty"`
    // 空气悬挂
    AirSuspension   string `json:"air_suspension,omitempty" xml:"air_suspension,omitempty"`
    // 车内氛围灯
    AmbientesLamp   string `json:"ambientes_lamp,omitempty" xml:"ambientes_lamp,omitempty"`
    // 车窗防夹手功能
    AntiPinchGlass   string `json:"anti_pinch_glass,omitempty" xml:"anti_pinch_glass,omitempty"`
    // 牵引力控制（ASR/TCS/TRC等）
    Asr   string `json:"asr,omitempty" xml:"asr,omitempty"`
    // 自动空调
    AutoAc   string `json:"auto_ac,omitempty" xml:"auto_ac,omitempty"`
    // 后视镜自动防眩目
    AutoDimmingInsideMirror   string `json:"auto_dimming_inside_mirror,omitempty" xml:"auto_dimming_inside_mirror,omitempty"`
    // 自动泊车入位
    AutomaticVarking   string `json:"automatic_varking,omitempty" xml:"automatic_varking,omitempty"`
    // nightVision
    Blis   string `json:"blis,omitempty" xml:"blis,omitempty"`
    // 蓝牙车载电话
    Bluetooth   string `json:"bluetooth,omitempty" xml:"bluetooth,omitempty"`
    // 车体结构
    BodyStructure   string `json:"body_structure,omitempty" xml:"body_structure,omitempty"`
    // 车身型式
    BodyType   string `json:"body_type,omitempty" xml:"body_type,omitempty"`
    // 缸径
    Bore   string `json:"bore,omitempty" xml:"bore,omitempty"`
    // 品牌代码
    BrandCode   string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
    // 品牌名称
    BrandName   string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
    // 车篷型式
    Calash   string `json:"calash,omitempty" xml:"calash,omitempty"`
    // 三元催化器
    Catalyst   string `json:"catalyst,omitempty" xml:"catalyst,omitempty"`
    // 中控锁
    CentralLocking   string `json:"central_locking,omitempty" xml:"central_locking,omitempty"`
    // 压缩底盘id
    ChassisCid   string `json:"chassis_cid,omitempty" xml:"chassis_cid,omitempty"`
    // 底盘号
    ChassisCode   string `json:"chassis_code,omitempty" xml:"chassis_code,omitempty"`
    // 工信部综合油耗（L/100KM）
    CombinedFuelConsumption   string `json:"combined_fuel_consumption,omitempty" xml:"combined_fuel_consumption,omitempty"`
    // 压缩比
    CompressionRatio   string `json:"compression_ratio,omitempty" xml:"compression_ratio,omitempty"`
    // 冷却方式
    CoolingMethod   string `json:"cooling_method,omitempty" xml:"cooling_method,omitempty"`
    // 转向头灯
    CornerHeadlamp   string `json:"corner_headlamp,omitempty" xml:"corner_headlamp,omitempty"`
    // 国别
    Country   string `json:"country,omitempty" xml:"country,omitempty"`
    // 定速巡航
    CruiseControl   string `json:"cruise_control,omitempty" xml:"cruise_control,omitempty"`
    // 整备质量（kg）
    CurbWeight   string `json:"curb_weight,omitempty" xml:"curb_weight,omitempty"`
    // 气缸排列形式
    CylinderArrangement   string `json:"cylinder_arrangement,omitempty" xml:"cylinder_arrangement,omitempty"`
    // 缸体材料
    CylinderBlock   string `json:"cylinder_block,omitempty" xml:"cylinder_block,omitempty"`
    // 缸盖材料
    CylinderHead   string `json:"cylinder_head,omitempty" xml:"cylinder_head,omitempty"`
    // 气缸容积
    CylinderVolume   string `json:"cylinder_volume,omitempty" xml:"cylinder_volume,omitempty"`
    // 气缸数（个）
    Cylinders   string `json:"cylinders,omitempty" xml:"cylinders,omitempty"`
    // 日间行车灯
    DaytimeRunningLamp   string `json:"daytime_running_lamp,omitempty" xml:"daytime_running_lamp,omitempty"`
    // 接近角
    DepartureAngle   string `json:"departure_angle,omitempty" xml:"departure_angle,omitempty"`
    // 近光灯
    DippedLights   string `json:"dipped_lights,omitempty" xml:"dipped_lights,omitempty"`
    // 排量
    Displacement   string `json:"displacement,omitempty" xml:"displacement,omitempty"`
    // 车门数
    Doors   string `json:"doors,omitempty" xml:"doors,omitempty"`
    // 驱动方式
    DriveMode   string `json:"drive_mode,omitempty" xml:"drive_mode,omitempty"`
    // 驱动形式
    DriveModel   string `json:"drive_model,omitempty" xml:"drive_model,omitempty"`
    // 驾驶座安全气囊
    DriverAirbag   string `json:"driver_airbag,omitempty" xml:"driver_airbag,omitempty"`
    // 驾驶座座椅电动调节
    DriverSeatPowerAdjustable   string `json:"driver_seat_power_adjustable,omitempty" xml:"driver_seat_power_adjustable,omitempty"`
    // 刹车辅助（EBA/BAS/BA等）
    Eba   string `json:"eba,omitempty" xml:"eba,omitempty"`
    // 制动力分配（EBD/CBC等）
    Ebd   string `json:"ebd,omitempty" xml:"ebd,omitempty"`
    // 方向盘电动调节
    ElectricAdjustableSteeringWheel   string `json:"electric_adjustable_steering_wheel,omitempty" xml:"electric_adjustable_steering_wheel,omitempty"`
    // 电动吸合门
    ElectricSuctionDoor   string `json:"electric_suction_door,omitempty" xml:"electric_suction_door,omitempty"`
    // 后视镜电动调节
    ElectricallyAdjustableOutsideMirror   string `json:"electrically_adjustable_outside_mirror,omitempty" xml:"electrically_adjustable_outside_mirror,omitempty"`
    // 排放标准
    EmissionStandard   string `json:"emission_standard,omitempty" xml:"emission_standard,omitempty"`
    // 发动机电子防盗
    EngineAntitheft   string `json:"engine_antitheft,omitempty" xml:"engine_antitheft,omitempty"`
    // 发动机描述
    EngineDescription   string `json:"engine_description,omitempty" xml:"engine_description,omitempty"`
    // 发动机特有技术
    EngineKnowhow   string `json:"engine_knowhow,omitempty" xml:"engine_knowhow,omitempty"`
    // 发动机位置
    EngineLocation   string `json:"engine_location,omitempty" xml:"engine_location,omitempty"`
    // 发动机生产厂家
    EngineManufacturer   string `json:"engine_manufacturer,omitempty" xml:"engine_manufacturer,omitempty"`
    // 发动机型号
    EngineModel   string `json:"engine_model,omitempty" xml:"engine_model,omitempty"`
    // 外接音源接口（AUX/USB/iPod等）
    EntertainmentConnector   string `json:"entertainment_connector,omitempty" xml:"entertainment_connector,omitempty"`
    // 自动驻车上坡辅助
    Epb   string `json:"epb,omitempty" xml:"epb,omitempty"`
    // 车身稳定控制（ESP/DSC/VSC等）
    Esp   string `json:"esp,omitempty" xml:"esp,omitempty"`
    // 前制动器类型
    FrontBrake   string `json:"front_brake,omitempty" xml:"front_brake,omitempty"`
    // 前座中央扶手
    FrontCenterArmrest   string `json:"front_center_armrest,omitempty" xml:"front_center_armrest,omitempty"`
    // 前排头部气囊（气帘）
    FrontCurtainAirbag   string `json:"front_curtain_airbag,omitempty" xml:"front_curtain_airbag,omitempty"`
    // 前雾灯
    FrontFogLamp   string `json:"front_fog_lamp,omitempty" xml:"front_fog_lamp,omitempty"`
    // 前电动车窗
    FrontPowerWindow   string `json:"front_power_window,omitempty" xml:"front_power_window,omitempty"`
    // 前轮毂规格
    FrontRim   string `json:"front_rim,omitempty" xml:"front_rim,omitempty"`
    // 前排座椅加热
    FrontSeatHeater   string `json:"front_seat_heater,omitempty" xml:"front_seat_heater,omitempty"`
    // 前排侧气囊
    FrontSideAirbag   string `json:"front_side_airbag,omitempty" xml:"front_side_airbag,omitempty"`
    // 前悬挂类型
    FrontSuspension   string `json:"front_suspension,omitempty" xml:"front_suspension,omitempty"`
    // 前轮距（mm）
    FrontTrack   string `json:"front_track,omitempty" xml:"front_track,omitempty"`
    // 前轮胎规格
    FrontTyre   string `json:"front_tyre,omitempty" xml:"front_tyre,omitempty"`
    // 燃油标号
    FuelGrade   string `json:"fuel_grade,omitempty" xml:"fuel_grade,omitempty"`
    // 供油方式
    FuelInjection   string `json:"fuel_injection,omitempty" xml:"fuel_injection,omitempty"`
    // 油箱容积（L）
    FuelTankCapacity   string `json:"fuel_tank_capacity,omitempty" xml:"fuel_tank_capacity,omitempty"`
    // 燃油类型
    FuelType   string `json:"fuel_type,omitempty" xml:"fuel_type,omitempty"`
    // 档位数
    GearNumber   string `json:"gear_number,omitempty" xml:"gear_number,omitempty"`
    // 代数
    Generation   string `json:"generation,omitempty" xml:"generation,omitempty"`
    // GPS导航
    Gps   string `json:"gps,omitempty" xml:"gps,omitempty"`
    // 指导价格(万元)
    GuidingPrice   string `json:"guiding_price,omitempty" xml:"guiding_price,omitempty"`
    // 陡坡缓降
    Hdc   string `json:"hdc,omitempty" xml:"hdc,omitempty"`
    // 大灯清洗装置
    HeadlampWasher   string `json:"headlamp_washer,omitempty" xml:"headlamp_washer,omitempty"`
    // 后视镜加热
    HeatedOutsideMirror   string `json:"heated_outside_mirror,omitempty" xml:"heated_outside_mirror,omitempty"`
    // 高度（mm）
    Height   string `json:"height,omitempty" xml:"height,omitempty"`
    // 大灯高度可调
    HeightAdjustableHeadlamp   string `json:"height_adjustable_headlamp,omitempty" xml:"height_adjustable_headlamp,omitempty"`
    // 座椅高低调节
    HeightAdjustableSeat   string `json:"height_adjustable_seat,omitempty" xml:"height_adjustable_seat,omitempty"`
    // 方向盘上下调节
    HeightAdjustableSteeringWheel   string `json:"height_adjustable_steering_wheel,omitempty" xml:"height_adjustable_steering_wheel,omitempty"`
    // 氙气大灯
    HidHeadlamp   string `json:"hid_headlamp,omitempty" xml:"hid_headlamp,omitempty"`
    // 远光灯
    HighBeam   string `json:"high_beam,omitempty" xml:"high_beam,omitempty"`
    // 最大马力（PS）
    Horsepower   string `json:"horsepower,omitempty" xml:"horsepower,omitempty"`
    // HUD抬头数字显示
    Hud   string `json:"hud,omitempty" xml:"hud,omitempty"`
    // 停产年份
    IdlingYear   string `json:"idling_year,omitempty" xml:"idling_year,omitempty"`
    // 进气形式
    Induction   string `json:"induction,omitempty" xml:"induction,omitempty"`
    // 后视镜记忆
    InsideMirrorWithMemory   string `json:"inside_mirror_with_memory,omitempty" xml:"inside_mirror_with_memory,omitempty"`
    // 隔热玻璃
    InsulatedGlass   string `json:"insulated_glass,omitempty" xml:"insulated_glass,omitempty"`
    // 发动机启停技术
    IntelligentStopStart   string `json:"intelligent_stop_start,omitempty" xml:"intelligent_stop_start,omitempty"`
    // 定位互动服务
    InteractiveInformationServices   string `json:"interactive_information_services,omitempty" xml:"interactive_information_services,omitempty"`
    // 内置硬盘
    InternalHardDisk   string `json:"internal_hard_disk,omitempty" xml:"internal_hard_disk,omitempty"`
    // 是否商用车
    IsCommercial   string `json:"is_commercial,omitempty" xml:"is_commercial,omitempty"`
    // 是否海外车
    IsForeign   string `json:"is_foreign,omitempty" xml:"is_foreign,omitempty"`
    // ISOFIX儿童座椅接口
    Isofix   string `json:"isofix,omitempty" xml:"isofix,omitempty"`
    // 无钥匙进入系统
    KeylessEntry   string `json:"keyless_entry,omitempty" xml:"keyless_entry,omitempty"`
    // 无钥匙启动系统
    KeylessGo   string `json:"keyless_go,omitempty" xml:"keyless_go,omitempty"`
    // 膝部气囊
    KneeAirbag   string `json:"knee_airbag,omitempty" xml:"knee_airbag,omitempty"`
    // LATCH座椅接口
    Latch   string `json:"latch,omitempty" xml:"latch,omitempty"`
    // 中控台彩色大屏
    LcdScreen   string `json:"lcd_screen,omitempty" xml:"lcd_screen,omitempty"`
    // 真皮座椅
    LeatherSeat   string `json:"leather_seat,omitempty" xml:"leather_seat,omitempty"`
    // 真皮方向盘
    LeatherSteeringWheel   string `json:"leather_steering_wheel,omitempty" xml:"leather_steering_wheel,omitempty"`
    // LED大灯
    LedHeadlamp   string `json:"led_headlamp,omitempty" xml:"led_headlamp,omitempty"`
    // 长度（mm）
    Length   string `json:"length,omitempty" xml:"length,omitempty"`
    // 方向盘前后调节
    LengthAdjustableSteeringWheel   string `json:"length_adjustable_steering_wheel,omitempty" xml:"length_adjustable_steering_wheel,omitempty"`
    // 车系
    LineName   string `json:"line_name,omitempty" xml:"line_name,omitempty"`
    // 上市月份
    ListingMonth   string `json:"listing_month,omitempty" xml:"listing_month,omitempty"`
    // 上市年份
    ListingYear   string `json:"listing_year,omitempty" xml:"listing_year,omitempty"`
    // 行李厢容积（L）
    LuggagePlace   string `json:"luggage_place,omitempty" xml:"luggage_place,omitempty"`
    // 腰部支撑调节
    LumberSupportAdjustable   string `json:"lumber_support_adjustable,omitempty" xml:"lumber_support_adjustable,omitempty"`
    // 力洋id
    LyId   string `json:"ly_id,omitempty" xml:"ly_id,omitempty"`
    // 人机交互系统
    ManMachineInteractiveSystem   string `json:"man_machine_interactive_system,omitempty" xml:"man_machine_interactive_system,omitempty"`
    // 厂家名称
    ManufactureName   string `json:"manufacture_name,omitempty" xml:"manufacture_name,omitempty"`
    // 厂家代码
    ManufacturersCode   string `json:"manufacturers_code,omitempty" xml:"manufacturers_code,omitempty"`
    // 最大载重质量（kg）
    MaxLoading   string `json:"max_loading,omitempty" xml:"max_loading,omitempty"`
    // 最高车速（km/h）
    MaxSpeed   string `json:"max_speed,omitempty" xml:"max_speed,omitempty"`
    // 电动座椅记忆
    MemorySeat   string `json:"memory_seat,omitempty" xml:"memory_seat,omitempty"`
    // 最小离地间隙（mm）
    MinGroundClearance   string `json:"min_ground_clearance,omitempty" xml:"min_ground_clearance,omitempty"`
    // 最小转弯半径
    MinTurningRadius   string `json:"min_turning_radius,omitempty" xml:"min_turning_radius,omitempty"`
    // 工信部车型代码
    ModelCode   string `json:"model_code,omitempty" xml:"model_code,omitempty"`
    // 车型
    ModelName   string `json:"model_name,omitempty" xml:"model_name,omitempty"`
    // 车型代码
    ModelsCode   string `json:"models_code,omitempty" xml:"models_code,omitempty"`
    // 音频支持MP3
    Mp3   string `json:"mp3,omitempty" xml:"mp3,omitempty"`
    // 多碟CD
    MultiDiscCd   string `json:"multi_disc_cd,omitempty" xml:"multi_disc_cd,omitempty"`
    // 多碟DVD
    MultiDiscDvd   string `json:"multi_disc_dvd,omitempty" xml:"multi_disc_dvd,omitempty"`
    // 多功能方向盘
    MultifunctionSteeringWheel   string `json:"multifunction_steering_wheel,omitempty" xml:"multifunction_steering_wheel,omitempty"`
    // 夜视系统
    NightVision   string `json:"night_vision,omitempty" xml:"night_vision,omitempty"`
    // 后排座椅整体放倒
    OverallRearSeatsFoldDown   string `json:"overall_rear_seats_fold_down,omitempty" xml:"overall_rear_seats_fold_down,omitempty"`
    // 全景摄像头
    PanoramicCamera   string `json:"panoramic_camera,omitempty" xml:"panoramic_camera,omitempty"`
    // 全景天窗
    PanoramicSunroof   string `json:"panoramic_sunroof,omitempty" xml:"panoramic_sunroof,omitempty"`
    // 泊车辅助
    ParkingAssist   string `json:"parking_assist,omitempty" xml:"parking_assist,omitempty"`
    // 驻车制动类型
    ParkingBrake   string `json:"parking_brake,omitempty" xml:"parking_brake,omitempty"`
    // 副驾驶安全气囊
    PassengerAirbag   string `json:"passenger_airbag,omitempty" xml:"passenger_airbag,omitempty"`
    // 副驾驶座座椅电动调节
    PassengerSeatPowerAdjustable   string `json:"passenger_seat_power_adjustable,omitempty" xml:"passenger_seat_power_adjustable,omitempty"`
    // 空气调节/花粉过滤
    PollenFilter   string `json:"pollen_filter,omitempty" xml:"pollen_filter,omitempty"`
    // 后视镜电动折叠
    PowerFoldOutsideMirror   string `json:"power_fold_outside_mirror,omitempty" xml:"power_fold_outside_mirror,omitempty"`
    // 最大功率（kW）
    PowerKw   string `json:"power_kw,omitempty" xml:"power_kw,omitempty"`
    // 最大功率转速（rpm）
    PowerRpm   string `json:"power_rpm,omitempty" xml:"power_rpm,omitempty"`
    // 助力类型
    PowerSteering   string `json:"power_steering,omitempty" xml:"power_steering,omitempty"`
    // 电动后备箱
    PowerTailgate   string `json:"power_tailgate,omitempty" xml:"power_tailgate,omitempty"`
    // 生产年份
    ProducedYear   string `json:"produced_year,omitempty" xml:"produced_year,omitempty"`
    // 生产状态
    ProductionStatus   string `json:"production_status,omitempty" xml:"production_status,omitempty"`
    // 感应雨刷
    RainSensingWipers   string `json:"rain_sensing_wipers,omitempty" xml:"rain_sensing_wipers,omitempty"`
    // 后排独立空调
    RearAc   string `json:"rear_ac,omitempty" xml:"rear_ac,omitempty"`
    // 后风挡遮阳帘
    RearBackWindowGlassBlind   string `json:"rear_back_window_glass_blind,omitempty" xml:"rear_back_window_glass_blind,omitempty"`
    // 后制动器类型
    RearBrake   string `json:"rear_brake,omitempty" xml:"rear_brake,omitempty"`
    // 后座中央扶手
    RearCenterArmrest   string `json:"rear_center_armrest,omitempty" xml:"rear_center_armrest,omitempty"`
    // 后排杯架
    RearCupHolder   string `json:"rear_cup_holder,omitempty" xml:"rear_cup_holder,omitempty"`
    // 后排头部气囊（气帘）
    RearCurtainAirbag   string `json:"rear_curtain_airbag,omitempty" xml:"rear_curtain_airbag,omitempty"`
    // 后排液晶屏
    RearEntertainmentScreen   string `json:"rear_entertainment_screen,omitempty" xml:"rear_entertainment_screen,omitempty"`
    // 倒车雷达
    RearParkingAid   string `json:"rear_parking_aid,omitempty" xml:"rear_parking_aid,omitempty"`
    // 后电动车窗
    RearPowerWindow   string `json:"rear_power_window,omitempty" xml:"rear_power_window,omitempty"`
    // 后轮毂规格
    RearRim   string `json:"rear_rim,omitempty" xml:"rear_rim,omitempty"`
    // 后排座椅加热
    RearSeatHeater   string `json:"rear_seat_heater,omitempty" xml:"rear_seat_heater,omitempty"`
    // 后排座椅电动调节
    RearSeatPowerAdjustable   string `json:"rear_seat_power_adjustable,omitempty" xml:"rear_seat_power_adjustable,omitempty"`
    // 后排座椅比例放倒
    RearSeatsProportionFoldDown   string `json:"rear_seats_proportion_fold_down,omitempty" xml:"rear_seats_proportion_fold_down,omitempty"`
    // 后排侧气囊
    RearSideAirbag   string `json:"rear_side_airbag,omitempty" xml:"rear_side_airbag,omitempty"`
    // 后排侧遮阳帘
    RearSideWindowGlassBlind   string `json:"rear_side_window_glass_blind,omitempty" xml:"rear_side_window_glass_blind,omitempty"`
    // 后悬挂类型
    RearSuspension   string `json:"rear_suspension,omitempty" xml:"rear_suspension,omitempty"`
    // 后轮距（mm）
    RearTrack   string `json:"rear_track,omitempty" xml:"rear_track,omitempty"`
    // 后轮胎规格
    RearTyre   string `json:"rear_tyre,omitempty" xml:"rear_tyre,omitempty"`
    // 后座出风口
    RearVent   string `json:"rear_vent,omitempty" xml:"rear_vent,omitempty"`
    // 倒车视频影像
    RearViewCamera   string `json:"rear_view_camera,omitempty" xml:"rear_view_camera,omitempty"`
    // 后雨刷
    RearWiper   string `json:"rear_wiper,omitempty" xml:"rear_wiper,omitempty"`
    // 车载冰箱
    Refrigerator   string `json:"refrigerator,omitempty" xml:"refrigerator,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 遥控钥匙
    RemoteControl   string `json:"remote_control,omitempty" xml:"remote_control,omitempty"`
    // 轮毂材料
    RimsMaterial   string `json:"rims_material,omitempty" xml:"rims_material,omitempty"`
    // 车顶型式
    RoofType   string `json:"roof_type,omitempty" xml:"roof_type,omitempty"`
    // 零胎压继续行驶
    RunFlatTyre   string `json:"run_flat_tyre,omitempty" xml:"run_flat_tyre,omitempty"`
    // 销售名称
    SalesName   string `json:"sales_name,omitempty" xml:"sales_name,omitempty"`
    // 销售状态
    SalesStatus   string `json:"sales_status,omitempty" xml:"sales_status,omitempty"`
    // 销售版本
    SalesVersion   string `json:"sales_version,omitempty" xml:"sales_version,omitempty"`
    // 年款
    SalesYear   string `json:"sales_year,omitempty" xml:"sales_year,omitempty"`
    // 座椅按摩
    SeatMassage   string `json:"seat_massage,omitempty" xml:"seat_massage,omitempty"`
    // 座椅通风
    SeatVentilation   string `json:"seat_ventilation,omitempty" xml:"seat_ventilation,omitempty"`
    // 安全带未系提示
    SeatbeltWarningLamp   string `json:"seatbelt_warning_lamp,omitempty" xml:"seatbelt_warning_lamp,omitempty"`
    // 座位数（个）
    Seats   string `json:"seats,omitempty" xml:"seats,omitempty"`
    // 第二排靠背角度调节
    SecondRowBackrestAdjustable   string `json:"second_row_backrest_adjustable,omitempty" xml:"second_row_backrest_adjustable,omitempty"`
    // 第二排座椅移动
    SecondRowSeatPositionAdjustable   string `json:"second_row_seat_position_adjustable,omitempty" xml:"second_row_seat_position_adjustable,omitempty"`
    // 车系代码
    SeriesCode   string `json:"series_code,omitempty" xml:"series_code,omitempty"`
    // 肩部支撑调节
    ShoulderSupportAdjustable   string `json:"shoulder_support_adjustable,omitempty" xml:"shoulder_support_adjustable,omitempty"`
    // 单碟CD
    SingleDiscCd   string `json:"single_disc_cd,omitempty" xml:"single_disc_cd,omitempty"`
    // 单碟DVD
    SingleDiscDvd   string `json:"single_disc_dvd,omitempty" xml:"single_disc_dvd,omitempty"`
    // 备胎规格
    SpareWheel   string `json:"spare_wheel,omitempty" xml:"spare_wheel,omitempty"`
    // 扬声器数量
    SpeakerNumber   string `json:"speaker_number,omitempty" xml:"speaker_number,omitempty"`
    // 中控液晶屏分屏显示
    Splitview   string `json:"splitview,omitempty" xml:"splitview,omitempty"`
    // 运动外观套件
    SportBodyDressUpKits   string `json:"sport_body_dress_up_kits,omitempty" xml:"sport_body_dress_up_kits,omitempty"`
    // 运动座椅
    SportSeat   string `json:"sport_seat,omitempty" xml:"sport_seat,omitempty"`
    // 转向机形式
    Steering   string `json:"steering,omitempty" xml:"steering,omitempty"`
    // 方向盘换挡
    SteeringWheelWithShift   string `json:"steering_wheel_with_shift,omitempty" xml:"steering_wheel_with_shift,omitempty"`
    // 冲程
    Stroke   string `json:"stroke,omitempty" xml:"stroke,omitempty"`
    // 市郊工况油耗（L/100KM）
    SuburbFuelConsumption   string `json:"suburb_fuel_consumption,omitempty" xml:"suburb_fuel_consumption,omitempty"`
    // 电动天窗
    Sunroof   string `json:"sunroof,omitempty" xml:"sunroof,omitempty"`
    // 遮阳板化妆镜
    SunvisorMirror   string `json:"sunvisor_mirror,omitempty" xml:"sunvisor_mirror,omitempty"`
    // 车载信息服务
    Telematics   string `json:"telematics,omitempty" xml:"telematics,omitempty"`
    // 第三排座椅
    ThirdRowSeat   string `json:"third_row_seat,omitempty" xml:"third_row_seat,omitempty"`
    // 胎压监测装置
    TirePressureMonitor   string `json:"tire_pressure_monitor,omitempty" xml:"tire_pressure_monitor,omitempty"`
    // 最大扭矩（N·m）
    TorqueNm   string `json:"torque_nm,omitempty" xml:"torque_nm,omitempty"`
    // 最大扭矩转速（rpm）
    TorqueRpm   string `json:"torque_rpm,omitempty" xml:"torque_rpm,omitempty"`
    // 变速器配置
    TransmissionConfiguration   string `json:"transmission_configuration,omitempty" xml:"transmission_configuration,omitempty"`
    // 变速器描述
    TransmissionDescription   string `json:"transmission_description,omitempty" xml:"transmission_description,omitempty"`
    // 变速器型号
    TransmissionModel   string `json:"transmission_model,omitempty" xml:"transmission_model,omitempty"`
    // 变速器类型
    TransmissionType   string `json:"transmission_type,omitempty" xml:"transmission_type,omitempty"`
    // 行车电脑显示屏
    TripComputer   string `json:"trip_computer,omitempty" xml:"trip_computer,omitempty"`
    // 市区工况油耗（L/100KM）
    UrbanFuelConsumption   string `json:"urban_fuel_consumption,omitempty" xml:"urban_fuel_consumption,omitempty"`
    // 配气机构
    ValveSystem   string `json:"valve_system,omitempty" xml:"valve_system,omitempty"`
    // 每缸气门数（个）
    ValvesPerCylinder   string `json:"valves_per_cylinder,omitempty" xml:"valves_per_cylinder,omitempty"`
    // 可变转向比
    VariableSteeringRatio   string `json:"variable_steering_ratio,omitempty" xml:"variable_steering_ratio,omitempty"`
    // 可变悬挂
    VariableSuspension   string `json:"variable_suspension,omitempty" xml:"variable_suspension,omitempty"`
    // 车身颜色
    VehicleColor   string `json:"vehicle_color,omitempty" xml:"vehicle_color,omitempty"`
    // 车辆级别
    VehicleSize   string `json:"vehicle_size,omitempty" xml:"vehicle_size,omitempty"`
    // 车载电视
    VehicleTv   string `json:"vehicle_tv,omitempty" xml:"vehicle_tv,omitempty"`
    // 车辆类型
    VehicleType   string `json:"vehicle_type,omitempty" xml:"vehicle_type,omitempty"`
    // 厂家类型（国产,合资,进口）
    VehicleVttributes   string `json:"vehicle_vttributes,omitempty" xml:"vehicle_vttributes,omitempty"`
    // 版本
    Version   string `json:"version,omitempty" xml:"version,omitempty"`
    // 压缩版本id
    VersionCid   string `json:"version_cid,omitempty" xml:"version_cid,omitempty"`
    // 虚拟多碟CD
    VirtualMultiDiscCd   string `json:"virtual_multi_disc_cd,omitempty" xml:"virtual_multi_disc_cd,omitempty"`
    // 整车质保
    WarrantyPeriod   string `json:"warranty_period,omitempty" xml:"warranty_period,omitempty"`
    // 轴距（mm）
    Wheelbase   string `json:"wheelbase,omitempty" xml:"wheelbase,omitempty"`
    // 宽度（mm）
    Width   string `json:"width,omitempty" xml:"width,omitempty"`
}
