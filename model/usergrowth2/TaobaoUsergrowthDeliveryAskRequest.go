package usergrowth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
广告投放询问 API请求
taobao.usergrowth.delivery.ask

提供给媒体在曝光广告前调用， 返回是否曝光以及曝光的物料信息
*/
type TaobaoUsergrowthDeliveryAskRequest struct {
    model.Params
    // 广告id，淘宝和媒体协商
    adid   string
    // 渠道标识，向淘宝技术申请
    channel   string
    // 是否爱折扣： 0： 不是? 1：是
    isDicountPrefer   string
    // 是否爱返现： 0： 不是? 1：是
    isCashPrefer   string
    // 用户网购支付层级：? 0： 0~50 1： 50~200?2： 200~500 3： 500以上
    payLayer   string
    // 是否爱赠品： 0： 不是? 1：是
    isGiftPrefer   string
    // 是否爱评价： 0： 不是? 1：是
    isRemarkUser   string
    // 是否爱分享： 0： 不是? 1：是
    isSharingUser   string
    // 宝宝性别，0:F 1:M 2:未知，但有宝宝 3: 没有宝宝
    babyGender   string
    // 是否有宠物?0：没有 1：有 2：预测有
    hasPet   string
    // 预测是否有房 0：没房 1：有房 2：预测有房
    hasHouse   string
    // 预测人生阶段
    lifeStage   string
    // 预测是否有车,0: 没有车 1：自己注册有车2：预测有车
    hasCar   string
    // 预测是否在校大学生?0：不是 1 是 2 预测是
    isUndergraduate   string
    // 预测月收入
    userIncome   string
    // 预测职业
    careerType   string
    // 预测学历,初高中,博士,专科,硕士,本科
    education   string
    // 注册IP解析的城市等级
    regCityLevel   string
    // 注册IP解析区县名
    ipAreaName   string
    // 注册IP解析城市名称
    ipCityName   string
    // 注册IP解析的省份名称
    ipProvName   string
    // 用户身份证号
    idCardNumber   string
    // 用户感兴趣的标签,多个用逗号隔开。如动漫、历史
    ukeywords   string
    // 年龄
    age   string
    // 性别, 0:未知 1： 男 2：女
    gender   string
    // 底价，单位为分
    adFloorPrice   string
    // 广告位支持图片格式 0: jpg, 1:?jpeg, 2: gif
    adImgType   string
    // 广告位位置 0:未知,1:头部,2:底部, 3:侧边栏,4:全屏 默认传0
    adPos   string
    // 广告位高度
    adHeight   string
    // 广告位宽度
    adWidth   string
    // 广告类型 0:横幅,1: 插屏, 2:开屏, 3:原生,4:视频
    adType   string
    // 设备经度
    geoLon   string
    // 设备维度
    geoLat   string
    // 设备屏幕纵向分辨率，单位：像素
    screenHeight   string
    // 设备屏幕水平分辨率，单位：像素
    screenWidth   string
    // 设备方向：0:未知， 1： 纵向；2： 横向
    orientation   string
    // 运营商，0: 未知， 1:移动,2:电信,3:联通
    carrier   string
    // 网络类型，0：未知，1：WIFI, 2: 2G, 3: 3G, 4: 4G; 5: 5G
    network   string
    // 手机品牌
    brand   string
    // 设备厂商
    made   string
    // 设备型号
    model   string
    // 素材展示所在的页面或者频道
    posCat   string
    // 合作方名称
    publishName   string
    // 给定appid中的位置id，用于提高转化率
    posId   string
    // app类别
    appCat   string
    // app版本号
    appVer   string
    // 对接的appName
    appName   string
    // 对接的appId
    appId   string
    // 关键词的描述
    description   string
    // 关键词
    keyword   string
    // 关键词类型
    keywordType   string
    // 安装转化目标app的标识， 0： 未安装转化目标app, 1: 安装了转化目标app
    appInstallFlag   string
    // 转化目标app, 1：淘宝；2：天猫；3：闲鱼；4：支付宝
    app   string
    // idfa的md5值， 32位小写
    idfaMd5   string
    // imei的md5值， 32位小写
    imeiMd5   string
    // idfa原生值
    idfa   string
    // imei原生值
    imei   string
    // 手机号
    mobile   string
    // 转化类型， 1： 激活；2： 新登；32896：定向促活
    transformType   string
    // 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
    os   string
    // 广告创意id，淘宝和媒体协商
    cid   string
    // android最新广告标识
    oaid   string
    // android最新广告标识md5值， 32位小写
    oaidMd5   string
}

// 初始化TaobaoUsergrowthDeliveryAskRequest对象
func NewTaobaoUsergrowthDeliveryAskRequest() *TaobaoUsergrowthDeliveryAskRequest{
    return &TaobaoUsergrowthDeliveryAskRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDeliveryAskRequest) GetApiMethodName() string {
    return "taobao.usergrowth.delivery.ask"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDeliveryAskRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Adid Setter
// 广告id，淘宝和媒体协商
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAdid(adid string) error {
    r.adid = adid
    r.Set("adid", adid)
    return nil
}

// Adid Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAdid() string {
    return r.adid
}
// Channel Setter
// 渠道标识，向淘宝技术申请
func (r *TaobaoUsergrowthDeliveryAskRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetChannel() string {
    return r.channel
}
// IsDicountPrefer Setter
// 是否爱折扣： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIsDicountPrefer(isDicountPrefer string) error {
    r.isDicountPrefer = isDicountPrefer
    r.Set("is_dicount_prefer", isDicountPrefer)
    return nil
}

// IsDicountPrefer Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIsDicountPrefer() string {
    return r.isDicountPrefer
}
// IsCashPrefer Setter
// 是否爱返现： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIsCashPrefer(isCashPrefer string) error {
    r.isCashPrefer = isCashPrefer
    r.Set("is_cash_prefer", isCashPrefer)
    return nil
}

// IsCashPrefer Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIsCashPrefer() string {
    return r.isCashPrefer
}
// PayLayer Setter
// 用户网购支付层级：? 0： 0~50 1： 50~200?2： 200~500 3： 500以上
func (r *TaobaoUsergrowthDeliveryAskRequest) SetPayLayer(payLayer string) error {
    r.payLayer = payLayer
    r.Set("pay_layer", payLayer)
    return nil
}

// PayLayer Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetPayLayer() string {
    return r.payLayer
}
// IsGiftPrefer Setter
// 是否爱赠品： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIsGiftPrefer(isGiftPrefer string) error {
    r.isGiftPrefer = isGiftPrefer
    r.Set("is_gift_prefer", isGiftPrefer)
    return nil
}

// IsGiftPrefer Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIsGiftPrefer() string {
    return r.isGiftPrefer
}
// IsRemarkUser Setter
// 是否爱评价： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIsRemarkUser(isRemarkUser string) error {
    r.isRemarkUser = isRemarkUser
    r.Set("is_remark_user", isRemarkUser)
    return nil
}

// IsRemarkUser Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIsRemarkUser() string {
    return r.isRemarkUser
}
// IsSharingUser Setter
// 是否爱分享： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIsSharingUser(isSharingUser string) error {
    r.isSharingUser = isSharingUser
    r.Set("is_sharing_user", isSharingUser)
    return nil
}

// IsSharingUser Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIsSharingUser() string {
    return r.isSharingUser
}
// BabyGender Setter
// 宝宝性别，0:F 1:M 2:未知，但有宝宝 3: 没有宝宝
func (r *TaobaoUsergrowthDeliveryAskRequest) SetBabyGender(babyGender string) error {
    r.babyGender = babyGender
    r.Set("baby_gender", babyGender)
    return nil
}

// BabyGender Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetBabyGender() string {
    return r.babyGender
}
// HasPet Setter
// 是否有宠物?0：没有 1：有 2：预测有
func (r *TaobaoUsergrowthDeliveryAskRequest) SetHasPet(hasPet string) error {
    r.hasPet = hasPet
    r.Set("has_pet", hasPet)
    return nil
}

// HasPet Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetHasPet() string {
    return r.hasPet
}
// HasHouse Setter
// 预测是否有房 0：没房 1：有房 2：预测有房
func (r *TaobaoUsergrowthDeliveryAskRequest) SetHasHouse(hasHouse string) error {
    r.hasHouse = hasHouse
    r.Set("has_house", hasHouse)
    return nil
}

// HasHouse Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetHasHouse() string {
    return r.hasHouse
}
// LifeStage Setter
// 预测人生阶段
func (r *TaobaoUsergrowthDeliveryAskRequest) SetLifeStage(lifeStage string) error {
    r.lifeStage = lifeStage
    r.Set("life_stage", lifeStage)
    return nil
}

// LifeStage Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetLifeStage() string {
    return r.lifeStage
}
// HasCar Setter
// 预测是否有车,0: 没有车 1：自己注册有车2：预测有车
func (r *TaobaoUsergrowthDeliveryAskRequest) SetHasCar(hasCar string) error {
    r.hasCar = hasCar
    r.Set("has_car", hasCar)
    return nil
}

// HasCar Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetHasCar() string {
    return r.hasCar
}
// IsUndergraduate Setter
// 预测是否在校大学生?0：不是 1 是 2 预测是
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIsUndergraduate(isUndergraduate string) error {
    r.isUndergraduate = isUndergraduate
    r.Set("is_undergraduate", isUndergraduate)
    return nil
}

// IsUndergraduate Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIsUndergraduate() string {
    return r.isUndergraduate
}
// UserIncome Setter
// 预测月收入
func (r *TaobaoUsergrowthDeliveryAskRequest) SetUserIncome(userIncome string) error {
    r.userIncome = userIncome
    r.Set("user_income", userIncome)
    return nil
}

// UserIncome Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetUserIncome() string {
    return r.userIncome
}
// CareerType Setter
// 预测职业
func (r *TaobaoUsergrowthDeliveryAskRequest) SetCareerType(careerType string) error {
    r.careerType = careerType
    r.Set("career_type", careerType)
    return nil
}

// CareerType Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetCareerType() string {
    return r.careerType
}
// Education Setter
// 预测学历,初高中,博士,专科,硕士,本科
func (r *TaobaoUsergrowthDeliveryAskRequest) SetEducation(education string) error {
    r.education = education
    r.Set("education", education)
    return nil
}

// Education Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetEducation() string {
    return r.education
}
// RegCityLevel Setter
// 注册IP解析的城市等级
func (r *TaobaoUsergrowthDeliveryAskRequest) SetRegCityLevel(regCityLevel string) error {
    r.regCityLevel = regCityLevel
    r.Set("reg_city_level", regCityLevel)
    return nil
}

// RegCityLevel Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetRegCityLevel() string {
    return r.regCityLevel
}
// IpAreaName Setter
// 注册IP解析区县名
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIpAreaName(ipAreaName string) error {
    r.ipAreaName = ipAreaName
    r.Set("ip_area_name", ipAreaName)
    return nil
}

// IpAreaName Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIpAreaName() string {
    return r.ipAreaName
}
// IpCityName Setter
// 注册IP解析城市名称
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIpCityName(ipCityName string) error {
    r.ipCityName = ipCityName
    r.Set("ip_city_name", ipCityName)
    return nil
}

// IpCityName Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIpCityName() string {
    return r.ipCityName
}
// IpProvName Setter
// 注册IP解析的省份名称
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIpProvName(ipProvName string) error {
    r.ipProvName = ipProvName
    r.Set("ip_prov_name", ipProvName)
    return nil
}

// IpProvName Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIpProvName() string {
    return r.ipProvName
}
// IdCardNumber Setter
// 用户身份证号
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIdCardNumber(idCardNumber string) error {
    r.idCardNumber = idCardNumber
    r.Set("id_card_number", idCardNumber)
    return nil
}

// IdCardNumber Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIdCardNumber() string {
    return r.idCardNumber
}
// Ukeywords Setter
// 用户感兴趣的标签,多个用逗号隔开。如动漫、历史
func (r *TaobaoUsergrowthDeliveryAskRequest) SetUkeywords(ukeywords string) error {
    r.ukeywords = ukeywords
    r.Set("ukeywords", ukeywords)
    return nil
}

// Ukeywords Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetUkeywords() string {
    return r.ukeywords
}
// Age Setter
// 年龄
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAge(age string) error {
    r.age = age
    r.Set("age", age)
    return nil
}

// Age Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAge() string {
    return r.age
}
// Gender Setter
// 性别, 0:未知 1： 男 2：女
func (r *TaobaoUsergrowthDeliveryAskRequest) SetGender(gender string) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

// Gender Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetGender() string {
    return r.gender
}
// AdFloorPrice Setter
// 底价，单位为分
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAdFloorPrice(adFloorPrice string) error {
    r.adFloorPrice = adFloorPrice
    r.Set("ad_floor_price", adFloorPrice)
    return nil
}

// AdFloorPrice Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAdFloorPrice() string {
    return r.adFloorPrice
}
// AdImgType Setter
// 广告位支持图片格式 0: jpg, 1:?jpeg, 2: gif
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAdImgType(adImgType string) error {
    r.adImgType = adImgType
    r.Set("ad_img_type", adImgType)
    return nil
}

// AdImgType Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAdImgType() string {
    return r.adImgType
}
// AdPos Setter
// 广告位位置 0:未知,1:头部,2:底部, 3:侧边栏,4:全屏 默认传0
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAdPos(adPos string) error {
    r.adPos = adPos
    r.Set("ad_pos", adPos)
    return nil
}

// AdPos Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAdPos() string {
    return r.adPos
}
// AdHeight Setter
// 广告位高度
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAdHeight(adHeight string) error {
    r.adHeight = adHeight
    r.Set("ad_height", adHeight)
    return nil
}

// AdHeight Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAdHeight() string {
    return r.adHeight
}
// AdWidth Setter
// 广告位宽度
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAdWidth(adWidth string) error {
    r.adWidth = adWidth
    r.Set("ad_width", adWidth)
    return nil
}

// AdWidth Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAdWidth() string {
    return r.adWidth
}
// AdType Setter
// 广告类型 0:横幅,1: 插屏, 2:开屏, 3:原生,4:视频
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAdType(adType string) error {
    r.adType = adType
    r.Set("ad_type", adType)
    return nil
}

// AdType Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAdType() string {
    return r.adType
}
// GeoLon Setter
// 设备经度
func (r *TaobaoUsergrowthDeliveryAskRequest) SetGeoLon(geoLon string) error {
    r.geoLon = geoLon
    r.Set("geo_lon", geoLon)
    return nil
}

// GeoLon Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetGeoLon() string {
    return r.geoLon
}
// GeoLat Setter
// 设备维度
func (r *TaobaoUsergrowthDeliveryAskRequest) SetGeoLat(geoLat string) error {
    r.geoLat = geoLat
    r.Set("geo_lat", geoLat)
    return nil
}

// GeoLat Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetGeoLat() string {
    return r.geoLat
}
// ScreenHeight Setter
// 设备屏幕纵向分辨率，单位：像素
func (r *TaobaoUsergrowthDeliveryAskRequest) SetScreenHeight(screenHeight string) error {
    r.screenHeight = screenHeight
    r.Set("screen_height", screenHeight)
    return nil
}

// ScreenHeight Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetScreenHeight() string {
    return r.screenHeight
}
// ScreenWidth Setter
// 设备屏幕水平分辨率，单位：像素
func (r *TaobaoUsergrowthDeliveryAskRequest) SetScreenWidth(screenWidth string) error {
    r.screenWidth = screenWidth
    r.Set("screen_width", screenWidth)
    return nil
}

// ScreenWidth Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetScreenWidth() string {
    return r.screenWidth
}
// Orientation Setter
// 设备方向：0:未知， 1： 纵向；2： 横向
func (r *TaobaoUsergrowthDeliveryAskRequest) SetOrientation(orientation string) error {
    r.orientation = orientation
    r.Set("orientation", orientation)
    return nil
}

// Orientation Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetOrientation() string {
    return r.orientation
}
// Carrier Setter
// 运营商，0: 未知， 1:移动,2:电信,3:联通
func (r *TaobaoUsergrowthDeliveryAskRequest) SetCarrier(carrier string) error {
    r.carrier = carrier
    r.Set("carrier", carrier)
    return nil
}

// Carrier Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetCarrier() string {
    return r.carrier
}
// Network Setter
// 网络类型，0：未知，1：WIFI, 2: 2G, 3: 3G, 4: 4G; 5: 5G
func (r *TaobaoUsergrowthDeliveryAskRequest) SetNetwork(network string) error {
    r.network = network
    r.Set("network", network)
    return nil
}

// Network Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetNetwork() string {
    return r.network
}
// Brand Setter
// 手机品牌
func (r *TaobaoUsergrowthDeliveryAskRequest) SetBrand(brand string) error {
    r.brand = brand
    r.Set("brand", brand)
    return nil
}

// Brand Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetBrand() string {
    return r.brand
}
// Made Setter
// 设备厂商
func (r *TaobaoUsergrowthDeliveryAskRequest) SetMade(made string) error {
    r.made = made
    r.Set("made", made)
    return nil
}

// Made Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetMade() string {
    return r.made
}
// Model Setter
// 设备型号
func (r *TaobaoUsergrowthDeliveryAskRequest) SetModel(model string) error {
    r.model = model
    r.Set("model", model)
    return nil
}

// Model Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetModel() string {
    return r.model
}
// PosCat Setter
// 素材展示所在的页面或者频道
func (r *TaobaoUsergrowthDeliveryAskRequest) SetPosCat(posCat string) error {
    r.posCat = posCat
    r.Set("pos_cat", posCat)
    return nil
}

// PosCat Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetPosCat() string {
    return r.posCat
}
// PublishName Setter
// 合作方名称
func (r *TaobaoUsergrowthDeliveryAskRequest) SetPublishName(publishName string) error {
    r.publishName = publishName
    r.Set("publish_name", publishName)
    return nil
}

// PublishName Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetPublishName() string {
    return r.publishName
}
// PosId Setter
// 给定appid中的位置id，用于提高转化率
func (r *TaobaoUsergrowthDeliveryAskRequest) SetPosId(posId string) error {
    r.posId = posId
    r.Set("pos_id", posId)
    return nil
}

// PosId Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetPosId() string {
    return r.posId
}
// AppCat Setter
// app类别
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAppCat(appCat string) error {
    r.appCat = appCat
    r.Set("app_cat", appCat)
    return nil
}

// AppCat Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAppCat() string {
    return r.appCat
}
// AppVer Setter
// app版本号
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAppVer(appVer string) error {
    r.appVer = appVer
    r.Set("app_ver", appVer)
    return nil
}

// AppVer Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAppVer() string {
    return r.appVer
}
// AppName Setter
// 对接的appName
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

// AppName Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAppName() string {
    return r.appName
}
// AppId Setter
// 对接的appId
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAppId(appId string) error {
    r.appId = appId
    r.Set("app_id", appId)
    return nil
}

// AppId Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAppId() string {
    return r.appId
}
// Description Setter
// 关键词的描述
func (r *TaobaoUsergrowthDeliveryAskRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetDescription() string {
    return r.description
}
// Keyword Setter
// 关键词
func (r *TaobaoUsergrowthDeliveryAskRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

// Keyword Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetKeyword() string {
    return r.keyword
}
// KeywordType Setter
// 关键词类型
func (r *TaobaoUsergrowthDeliveryAskRequest) SetKeywordType(keywordType string) error {
    r.keywordType = keywordType
    r.Set("keyword_type", keywordType)
    return nil
}

// KeywordType Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetKeywordType() string {
    return r.keywordType
}
// AppInstallFlag Setter
// 安装转化目标app的标识， 0： 未安装转化目标app, 1: 安装了转化目标app
func (r *TaobaoUsergrowthDeliveryAskRequest) SetAppInstallFlag(appInstallFlag string) error {
    r.appInstallFlag = appInstallFlag
    r.Set("app_install_flag", appInstallFlag)
    return nil
}

// AppInstallFlag Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetAppInstallFlag() string {
    return r.appInstallFlag
}
// App Setter
// 转化目标app, 1：淘宝；2：天猫；3：闲鱼；4：支付宝
func (r *TaobaoUsergrowthDeliveryAskRequest) SetApp(app string) error {
    r.app = app
    r.Set("app", app)
    return nil
}

// App Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetApp() string {
    return r.app
}
// IdfaMd5 Setter
// idfa的md5值， 32位小写
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIdfaMd5(idfaMd5 string) error {
    r.idfaMd5 = idfaMd5
    r.Set("idfa_md5", idfaMd5)
    return nil
}

// IdfaMd5 Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIdfaMd5() string {
    return r.idfaMd5
}
// ImeiMd5 Setter
// imei的md5值， 32位小写
func (r *TaobaoUsergrowthDeliveryAskRequest) SetImeiMd5(imeiMd5 string) error {
    r.imeiMd5 = imeiMd5
    r.Set("imei_md5", imeiMd5)
    return nil
}

// ImeiMd5 Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetImeiMd5() string {
    return r.imeiMd5
}
// Idfa Setter
// idfa原生值
func (r *TaobaoUsergrowthDeliveryAskRequest) SetIdfa(idfa string) error {
    r.idfa = idfa
    r.Set("idfa", idfa)
    return nil
}

// Idfa Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetIdfa() string {
    return r.idfa
}
// Imei Setter
// imei原生值
func (r *TaobaoUsergrowthDeliveryAskRequest) SetImei(imei string) error {
    r.imei = imei
    r.Set("imei", imei)
    return nil
}

// Imei Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetImei() string {
    return r.imei
}
// Mobile Setter
// 手机号
func (r *TaobaoUsergrowthDeliveryAskRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetMobile() string {
    return r.mobile
}
// TransformType Setter
// 转化类型， 1： 激活；2： 新登；32896：定向促活
func (r *TaobaoUsergrowthDeliveryAskRequest) SetTransformType(transformType string) error {
    r.transformType = transformType
    r.Set("transform_type", transformType)
    return nil
}

// TransformType Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetTransformType() string {
    return r.transformType
}
// Os Setter
// 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
func (r *TaobaoUsergrowthDeliveryAskRequest) SetOs(os string) error {
    r.os = os
    r.Set("os", os)
    return nil
}

// Os Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetOs() string {
    return r.os
}
// Cid Setter
// 广告创意id，淘宝和媒体协商
func (r *TaobaoUsergrowthDeliveryAskRequest) SetCid(cid string) error {
    r.cid = cid
    r.Set("cid", cid)
    return nil
}

// Cid Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetCid() string {
    return r.cid
}
// Oaid Setter
// android最新广告标识
func (r *TaobaoUsergrowthDeliveryAskRequest) SetOaid(oaid string) error {
    r.oaid = oaid
    r.Set("oaid", oaid)
    return nil
}

// Oaid Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetOaid() string {
    return r.oaid
}
// OaidMd5 Setter
// android最新广告标识md5值， 32位小写
func (r *TaobaoUsergrowthDeliveryAskRequest) SetOaidMd5(oaidMd5 string) error {
    r.oaidMd5 = oaidMd5
    r.Set("oaid_md5", oaidMd5)
    return nil
}

// OaidMd5 Getter
func (r TaobaoUsergrowthDeliveryAskRequest) GetOaidMd5() string {
    return r.oaidMd5
}
