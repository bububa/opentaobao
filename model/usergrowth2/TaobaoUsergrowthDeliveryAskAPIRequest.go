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
type TaobaoUsergrowthDeliveryAskAPIRequest struct {
    model.Params
    // 广告id，淘宝和媒体协商
    _adid   string
    // 渠道标识，向淘宝技术申请
    _channel   string
    // 是否爱折扣： 0： 不是? 1：是
    _isDicountPrefer   string
    // 是否爱返现： 0： 不是? 1：是
    _isCashPrefer   string
    // 用户网购支付层级：? 0： 0~50 1： 50~200?2： 200~500 3： 500以上
    _payLayer   string
    // 是否爱赠品： 0： 不是? 1：是
    _isGiftPrefer   string
    // 是否爱评价： 0： 不是? 1：是
    _isRemarkUser   string
    // 是否爱分享： 0： 不是? 1：是
    _isSharingUser   string
    // 宝宝性别，0:F 1:M 2:未知，但有宝宝 3: 没有宝宝
    _babyGender   string
    // 是否有宠物?0：没有 1：有 2：预测有
    _hasPet   string
    // 预测是否有房 0：没房 1：有房 2：预测有房
    _hasHouse   string
    // 预测人生阶段
    _lifeStage   string
    // 预测是否有车,0: 没有车 1：自己注册有车2：预测有车
    _hasCar   string
    // 预测是否在校大学生?0：不是 1 是 2 预测是
    _isUndergraduate   string
    // 预测月收入
    _userIncome   string
    // 预测职业
    _careerType   string
    // 预测学历,初高中,博士,专科,硕士,本科
    _education   string
    // 注册IP解析的城市等级
    _regCityLevel   string
    // 注册IP解析区县名
    _ipAreaName   string
    // 注册IP解析城市名称
    _ipCityName   string
    // 注册IP解析的省份名称
    _ipProvName   string
    // 用户身份证号
    _idCardNumber   string
    // 用户感兴趣的标签,多个用逗号隔开。如动漫、历史
    _ukeywords   string
    // 年龄
    _age   string
    // 性别, 0:未知 1： 男 2：女
    _gender   string
    // 底价，单位为分
    _adFloorPrice   string
    // 广告位支持图片格式 0: jpg, 1:?jpeg, 2: gif
    _adImgType   string
    // 广告位位置 0:未知,1:头部,2:底部, 3:侧边栏,4:全屏 默认传0
    _adPos   string
    // 广告位高度
    _adHeight   string
    // 广告位宽度
    _adWidth   string
    // 广告类型 0:横幅,1: 插屏, 2:开屏, 3:原生,4:视频
    _adType   string
    // 设备经度
    _geoLon   string
    // 设备维度
    _geoLat   string
    // 设备屏幕纵向分辨率，单位：像素
    _screenHeight   string
    // 设备屏幕水平分辨率，单位：像素
    _screenWidth   string
    // 设备方向：0:未知， 1： 纵向；2： 横向
    _orientation   string
    // 运营商，0: 未知， 1:移动,2:电信,3:联通
    _carrier   string
    // 网络类型，0：未知，1：WIFI, 2: 2G, 3: 3G, 4: 4G; 5: 5G
    _network   string
    // 手机品牌
    _brand   string
    // 设备厂商
    _made   string
    // 设备型号
    _model   string
    // 素材展示所在的页面或者频道
    _posCat   string
    // 合作方名称
    _publishName   string
    // 给定appid中的位置id，用于提高转化率
    _posId   string
    // app类别
    _appCat   string
    // app版本号
    _appVer   string
    // 对接的appName
    _appName   string
    // 对接的appId
    _appId   string
    // 关键词的描述
    _description   string
    // 关键词
    _keyword   string
    // 关键词类型
    _keywordType   string
    // 安装转化目标app的标识， 0： 未安装转化目标app, 1: 安装了转化目标app
    _appInstallFlag   string
    // 转化目标app, 1：淘宝；2：天猫；3：闲鱼；4：支付宝
    _app   string
    // idfa的md5值， 32位小写
    _idfaMd5   string
    // imei的md5值， 32位小写
    _imeiMd5   string
    // idfa原生值
    _idfa   string
    // imei原生值
    _imei   string
    // 手机号
    _mobile   string
    // 转化类型， 1： 激活；2： 新登；32896：定向促活
    _transformType   string
    // 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
    _os   string
    // 广告创意id，淘宝和媒体协商
    _cid   string
    // android最新广告标识
    _oaid   string
    // android最新广告标识md5值， 32位小写
    _oaidMd5   string
}

// 初始化TaobaoUsergrowthDeliveryAskAPIRequest对象
func NewTaobaoUsergrowthDeliveryAskRequest() *TaobaoUsergrowthDeliveryAskAPIRequest{
    return &TaobaoUsergrowthDeliveryAskAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetApiMethodName() string {
    return "taobao.usergrowth.delivery.ask"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Adid Setter
// 广告id，淘宝和媒体协商
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAdid(_adid string) error {
    r._adid = _adid
    r.Set("adid", _adid)
    return nil
}

// Adid Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAdid() string {
    return r._adid
}
// Channel Setter
// 渠道标识，向淘宝技术申请
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetChannel() string {
    return r._channel
}
// IsDicountPrefer Setter
// 是否爱折扣： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIsDicountPrefer(_isDicountPrefer string) error {
    r._isDicountPrefer = _isDicountPrefer
    r.Set("is_dicount_prefer", _isDicountPrefer)
    return nil
}

// IsDicountPrefer Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIsDicountPrefer() string {
    return r._isDicountPrefer
}
// IsCashPrefer Setter
// 是否爱返现： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIsCashPrefer(_isCashPrefer string) error {
    r._isCashPrefer = _isCashPrefer
    r.Set("is_cash_prefer", _isCashPrefer)
    return nil
}

// IsCashPrefer Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIsCashPrefer() string {
    return r._isCashPrefer
}
// PayLayer Setter
// 用户网购支付层级：? 0： 0~50 1： 50~200?2： 200~500 3： 500以上
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetPayLayer(_payLayer string) error {
    r._payLayer = _payLayer
    r.Set("pay_layer", _payLayer)
    return nil
}

// PayLayer Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetPayLayer() string {
    return r._payLayer
}
// IsGiftPrefer Setter
// 是否爱赠品： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIsGiftPrefer(_isGiftPrefer string) error {
    r._isGiftPrefer = _isGiftPrefer
    r.Set("is_gift_prefer", _isGiftPrefer)
    return nil
}

// IsGiftPrefer Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIsGiftPrefer() string {
    return r._isGiftPrefer
}
// IsRemarkUser Setter
// 是否爱评价： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIsRemarkUser(_isRemarkUser string) error {
    r._isRemarkUser = _isRemarkUser
    r.Set("is_remark_user", _isRemarkUser)
    return nil
}

// IsRemarkUser Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIsRemarkUser() string {
    return r._isRemarkUser
}
// IsSharingUser Setter
// 是否爱分享： 0： 不是? 1：是
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIsSharingUser(_isSharingUser string) error {
    r._isSharingUser = _isSharingUser
    r.Set("is_sharing_user", _isSharingUser)
    return nil
}

// IsSharingUser Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIsSharingUser() string {
    return r._isSharingUser
}
// BabyGender Setter
// 宝宝性别，0:F 1:M 2:未知，但有宝宝 3: 没有宝宝
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetBabyGender(_babyGender string) error {
    r._babyGender = _babyGender
    r.Set("baby_gender", _babyGender)
    return nil
}

// BabyGender Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetBabyGender() string {
    return r._babyGender
}
// HasPet Setter
// 是否有宠物?0：没有 1：有 2：预测有
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetHasPet(_hasPet string) error {
    r._hasPet = _hasPet
    r.Set("has_pet", _hasPet)
    return nil
}

// HasPet Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetHasPet() string {
    return r._hasPet
}
// HasHouse Setter
// 预测是否有房 0：没房 1：有房 2：预测有房
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetHasHouse(_hasHouse string) error {
    r._hasHouse = _hasHouse
    r.Set("has_house", _hasHouse)
    return nil
}

// HasHouse Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetHasHouse() string {
    return r._hasHouse
}
// LifeStage Setter
// 预测人生阶段
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetLifeStage(_lifeStage string) error {
    r._lifeStage = _lifeStage
    r.Set("life_stage", _lifeStage)
    return nil
}

// LifeStage Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetLifeStage() string {
    return r._lifeStage
}
// HasCar Setter
// 预测是否有车,0: 没有车 1：自己注册有车2：预测有车
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetHasCar(_hasCar string) error {
    r._hasCar = _hasCar
    r.Set("has_car", _hasCar)
    return nil
}

// HasCar Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetHasCar() string {
    return r._hasCar
}
// IsUndergraduate Setter
// 预测是否在校大学生?0：不是 1 是 2 预测是
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIsUndergraduate(_isUndergraduate string) error {
    r._isUndergraduate = _isUndergraduate
    r.Set("is_undergraduate", _isUndergraduate)
    return nil
}

// IsUndergraduate Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIsUndergraduate() string {
    return r._isUndergraduate
}
// UserIncome Setter
// 预测月收入
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetUserIncome(_userIncome string) error {
    r._userIncome = _userIncome
    r.Set("user_income", _userIncome)
    return nil
}

// UserIncome Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetUserIncome() string {
    return r._userIncome
}
// CareerType Setter
// 预测职业
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetCareerType(_careerType string) error {
    r._careerType = _careerType
    r.Set("career_type", _careerType)
    return nil
}

// CareerType Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetCareerType() string {
    return r._careerType
}
// Education Setter
// 预测学历,初高中,博士,专科,硕士,本科
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetEducation(_education string) error {
    r._education = _education
    r.Set("education", _education)
    return nil
}

// Education Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetEducation() string {
    return r._education
}
// RegCityLevel Setter
// 注册IP解析的城市等级
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetRegCityLevel(_regCityLevel string) error {
    r._regCityLevel = _regCityLevel
    r.Set("reg_city_level", _regCityLevel)
    return nil
}

// RegCityLevel Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetRegCityLevel() string {
    return r._regCityLevel
}
// IpAreaName Setter
// 注册IP解析区县名
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIpAreaName(_ipAreaName string) error {
    r._ipAreaName = _ipAreaName
    r.Set("ip_area_name", _ipAreaName)
    return nil
}

// IpAreaName Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIpAreaName() string {
    return r._ipAreaName
}
// IpCityName Setter
// 注册IP解析城市名称
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIpCityName(_ipCityName string) error {
    r._ipCityName = _ipCityName
    r.Set("ip_city_name", _ipCityName)
    return nil
}

// IpCityName Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIpCityName() string {
    return r._ipCityName
}
// IpProvName Setter
// 注册IP解析的省份名称
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIpProvName(_ipProvName string) error {
    r._ipProvName = _ipProvName
    r.Set("ip_prov_name", _ipProvName)
    return nil
}

// IpProvName Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIpProvName() string {
    return r._ipProvName
}
// IdCardNumber Setter
// 用户身份证号
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIdCardNumber(_idCardNumber string) error {
    r._idCardNumber = _idCardNumber
    r.Set("id_card_number", _idCardNumber)
    return nil
}

// IdCardNumber Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIdCardNumber() string {
    return r._idCardNumber
}
// Ukeywords Setter
// 用户感兴趣的标签,多个用逗号隔开。如动漫、历史
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetUkeywords(_ukeywords string) error {
    r._ukeywords = _ukeywords
    r.Set("ukeywords", _ukeywords)
    return nil
}

// Ukeywords Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetUkeywords() string {
    return r._ukeywords
}
// Age Setter
// 年龄
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAge(_age string) error {
    r._age = _age
    r.Set("age", _age)
    return nil
}

// Age Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAge() string {
    return r._age
}
// Gender Setter
// 性别, 0:未知 1： 男 2：女
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetGender(_gender string) error {
    r._gender = _gender
    r.Set("gender", _gender)
    return nil
}

// Gender Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetGender() string {
    return r._gender
}
// AdFloorPrice Setter
// 底价，单位为分
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAdFloorPrice(_adFloorPrice string) error {
    r._adFloorPrice = _adFloorPrice
    r.Set("ad_floor_price", _adFloorPrice)
    return nil
}

// AdFloorPrice Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAdFloorPrice() string {
    return r._adFloorPrice
}
// AdImgType Setter
// 广告位支持图片格式 0: jpg, 1:?jpeg, 2: gif
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAdImgType(_adImgType string) error {
    r._adImgType = _adImgType
    r.Set("ad_img_type", _adImgType)
    return nil
}

// AdImgType Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAdImgType() string {
    return r._adImgType
}
// AdPos Setter
// 广告位位置 0:未知,1:头部,2:底部, 3:侧边栏,4:全屏 默认传0
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAdPos(_adPos string) error {
    r._adPos = _adPos
    r.Set("ad_pos", _adPos)
    return nil
}

// AdPos Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAdPos() string {
    return r._adPos
}
// AdHeight Setter
// 广告位高度
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAdHeight(_adHeight string) error {
    r._adHeight = _adHeight
    r.Set("ad_height", _adHeight)
    return nil
}

// AdHeight Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAdHeight() string {
    return r._adHeight
}
// AdWidth Setter
// 广告位宽度
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAdWidth(_adWidth string) error {
    r._adWidth = _adWidth
    r.Set("ad_width", _adWidth)
    return nil
}

// AdWidth Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAdWidth() string {
    return r._adWidth
}
// AdType Setter
// 广告类型 0:横幅,1: 插屏, 2:开屏, 3:原生,4:视频
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAdType(_adType string) error {
    r._adType = _adType
    r.Set("ad_type", _adType)
    return nil
}

// AdType Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAdType() string {
    return r._adType
}
// GeoLon Setter
// 设备经度
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetGeoLon(_geoLon string) error {
    r._geoLon = _geoLon
    r.Set("geo_lon", _geoLon)
    return nil
}

// GeoLon Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetGeoLon() string {
    return r._geoLon
}
// GeoLat Setter
// 设备维度
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetGeoLat(_geoLat string) error {
    r._geoLat = _geoLat
    r.Set("geo_lat", _geoLat)
    return nil
}

// GeoLat Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetGeoLat() string {
    return r._geoLat
}
// ScreenHeight Setter
// 设备屏幕纵向分辨率，单位：像素
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetScreenHeight(_screenHeight string) error {
    r._screenHeight = _screenHeight
    r.Set("screen_height", _screenHeight)
    return nil
}

// ScreenHeight Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetScreenHeight() string {
    return r._screenHeight
}
// ScreenWidth Setter
// 设备屏幕水平分辨率，单位：像素
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetScreenWidth(_screenWidth string) error {
    r._screenWidth = _screenWidth
    r.Set("screen_width", _screenWidth)
    return nil
}

// ScreenWidth Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetScreenWidth() string {
    return r._screenWidth
}
// Orientation Setter
// 设备方向：0:未知， 1： 纵向；2： 横向
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetOrientation(_orientation string) error {
    r._orientation = _orientation
    r.Set("orientation", _orientation)
    return nil
}

// Orientation Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetOrientation() string {
    return r._orientation
}
// Carrier Setter
// 运营商，0: 未知， 1:移动,2:电信,3:联通
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetCarrier(_carrier string) error {
    r._carrier = _carrier
    r.Set("carrier", _carrier)
    return nil
}

// Carrier Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetCarrier() string {
    return r._carrier
}
// Network Setter
// 网络类型，0：未知，1：WIFI, 2: 2G, 3: 3G, 4: 4G; 5: 5G
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetNetwork(_network string) error {
    r._network = _network
    r.Set("network", _network)
    return nil
}

// Network Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetNetwork() string {
    return r._network
}
// Brand Setter
// 手机品牌
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetBrand(_brand string) error {
    r._brand = _brand
    r.Set("brand", _brand)
    return nil
}

// Brand Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetBrand() string {
    return r._brand
}
// Made Setter
// 设备厂商
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetMade(_made string) error {
    r._made = _made
    r.Set("made", _made)
    return nil
}

// Made Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetMade() string {
    return r._made
}
// Model Setter
// 设备型号
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetModel(_model string) error {
    r._model = _model
    r.Set("model", _model)
    return nil
}

// Model Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetModel() string {
    return r._model
}
// PosCat Setter
// 素材展示所在的页面或者频道
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetPosCat(_posCat string) error {
    r._posCat = _posCat
    r.Set("pos_cat", _posCat)
    return nil
}

// PosCat Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetPosCat() string {
    return r._posCat
}
// PublishName Setter
// 合作方名称
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetPublishName(_publishName string) error {
    r._publishName = _publishName
    r.Set("publish_name", _publishName)
    return nil
}

// PublishName Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetPublishName() string {
    return r._publishName
}
// PosId Setter
// 给定appid中的位置id，用于提高转化率
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetPosId(_posId string) error {
    r._posId = _posId
    r.Set("pos_id", _posId)
    return nil
}

// PosId Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetPosId() string {
    return r._posId
}
// AppCat Setter
// app类别
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAppCat(_appCat string) error {
    r._appCat = _appCat
    r.Set("app_cat", _appCat)
    return nil
}

// AppCat Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAppCat() string {
    return r._appCat
}
// AppVer Setter
// app版本号
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAppVer(_appVer string) error {
    r._appVer = _appVer
    r.Set("app_ver", _appVer)
    return nil
}

// AppVer Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAppVer() string {
    return r._appVer
}
// AppName Setter
// 对接的appName
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAppName() string {
    return r._appName
}
// AppId Setter
// 对接的appId
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAppId(_appId string) error {
    r._appId = _appId
    r.Set("app_id", _appId)
    return nil
}

// AppId Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAppId() string {
    return r._appId
}
// Description Setter
// 关键词的描述
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetDescription() string {
    return r._description
}
// Keyword Setter
// 关键词
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetKeyword() string {
    return r._keyword
}
// KeywordType Setter
// 关键词类型
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetKeywordType(_keywordType string) error {
    r._keywordType = _keywordType
    r.Set("keyword_type", _keywordType)
    return nil
}

// KeywordType Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetKeywordType() string {
    return r._keywordType
}
// AppInstallFlag Setter
// 安装转化目标app的标识， 0： 未安装转化目标app, 1: 安装了转化目标app
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetAppInstallFlag(_appInstallFlag string) error {
    r._appInstallFlag = _appInstallFlag
    r.Set("app_install_flag", _appInstallFlag)
    return nil
}

// AppInstallFlag Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetAppInstallFlag() string {
    return r._appInstallFlag
}
// App Setter
// 转化目标app, 1：淘宝；2：天猫；3：闲鱼；4：支付宝
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetApp(_app string) error {
    r._app = _app
    r.Set("app", _app)
    return nil
}

// App Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetApp() string {
    return r._app
}
// IdfaMd5 Setter
// idfa的md5值， 32位小写
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIdfaMd5(_idfaMd5 string) error {
    r._idfaMd5 = _idfaMd5
    r.Set("idfa_md5", _idfaMd5)
    return nil
}

// IdfaMd5 Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIdfaMd5() string {
    return r._idfaMd5
}
// ImeiMd5 Setter
// imei的md5值， 32位小写
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetImeiMd5(_imeiMd5 string) error {
    r._imeiMd5 = _imeiMd5
    r.Set("imei_md5", _imeiMd5)
    return nil
}

// ImeiMd5 Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetImeiMd5() string {
    return r._imeiMd5
}
// Idfa Setter
// idfa原生值
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetIdfa(_idfa string) error {
    r._idfa = _idfa
    r.Set("idfa", _idfa)
    return nil
}

// Idfa Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetIdfa() string {
    return r._idfa
}
// Imei Setter
// imei原生值
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetImei(_imei string) error {
    r._imei = _imei
    r.Set("imei", _imei)
    return nil
}

// Imei Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetImei() string {
    return r._imei
}
// Mobile Setter
// 手机号
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetMobile() string {
    return r._mobile
}
// TransformType Setter
// 转化类型， 1： 激活；2： 新登；32896：定向促活
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetTransformType(_transformType string) error {
    r._transformType = _transformType
    r.Set("transform_type", _transformType)
    return nil
}

// TransformType Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetTransformType() string {
    return r._transformType
}
// Os Setter
// 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetOs(_os string) error {
    r._os = _os
    r.Set("os", _os)
    return nil
}

// Os Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetOs() string {
    return r._os
}
// Cid Setter
// 广告创意id，淘宝和媒体协商
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetCid(_cid string) error {
    r._cid = _cid
    r.Set("cid", _cid)
    return nil
}

// Cid Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetCid() string {
    return r._cid
}
// Oaid Setter
// android最新广告标识
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetOaid(_oaid string) error {
    r._oaid = _oaid
    r.Set("oaid", _oaid)
    return nil
}

// Oaid Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetOaid() string {
    return r._oaid
}
// OaidMd5 Setter
// android最新广告标识md5值， 32位小写
func (r *TaobaoUsergrowthDeliveryAskAPIRequest) SetOaidMd5(_oaidMd5 string) error {
    r._oaidMd5 = _oaidMd5
    r.Set("oaid_md5", _oaidMd5)
    return nil
}

// OaidMd5 Getter
func (r TaobaoUsergrowthDeliveryAskAPIRequest) GetOaidMd5() string {
    return r._oaidMd5
}
