package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthDeliveryAskAPIRequest
广告投放询问 API请求
taobao.usergrowth.delivery.ask

提供给媒体在曝光广告前调用， 返回是否曝光以及曝光的物料信息 */
type TaobaoUsergrowthDeliveryAskAPIRequest struct {
	model.Params
	// 广告id，淘宝和媒体协商
	_adid string
	// 渠道标识，向淘宝技术申请
	_channel string
	// 是否爱折扣： 0： 不是? 1：是
	_isDicountPrefer string
	// 是否爱返现： 0： 不是? 1：是
	_isCashPrefer string
	// 用户网购支付层级：? 0： 0~50 1： 50~200?2： 200~500 3： 500以上
	_payLayer string
	// 是否爱赠品： 0： 不是? 1：是
	_isGiftPrefer string
	// 是否爱评价： 0： 不是? 1：是
	_isRemarkUser string
	// 是否爱分享： 0： 不是? 1：是
	_isSharingUser string
	// 宝宝性别，0:F 1:M 2:未知，但有宝宝 3: 没有宝宝
	_babyGender string
	// 是否有宠物?0：没有 1：有 2：预测有
	_hasPet string
	// 预测是否有房 0：没房 1：有房 2：预测有房
	_hasHouse string
	// 预测人生阶段
	_lifeStage string
	// 预测是否有车,0: 没有车 1：自己注册有车2：预测有车
	_hasCar string
	// 预测是否在校大学生?0：不是 1 是 2 预测是
	_isUndergraduate string
	// 预测月收入
	_userIncome string
	// 预测职业
	_careerType string
	// 预测学历,初高中,博士,专科,硕士,本科
	_education string
	// 注册IP解析的城市等级
	_regCityLevel string
	// 注册IP解析区县名
	_ipAreaName string
	// 注册IP解析城市名称
	_ipCityName string
	// 注册IP解析的省份名称
	_ipProvName string
	// 用户身份证号
	_idCardNumber string
	// 用户感兴趣的标签,多个用逗号隔开。如动漫、历史
	_ukeywords string
	// 年龄
	_age string
	// 性别, 0:未知 1： 男 2：女
	_gender string
	// 底价，单位为分
	_adFloorPrice string
	// 广告位支持图片格式 0: jpg, 1:?jpeg, 2: gif
	_adImgType string
	// 广告位位置 0:未知,1:头部,2:底部, 3:侧边栏,4:全屏 默认传0
	_adPos string
	// 广告位高度
	_adHeight string
	// 广告位宽度
	_adWidth string
	// 广告类型 0:横幅,1: 插屏, 2:开屏, 3:原生,4:视频
	_adType string
	// 设备经度
	_geoLon string
	// 设备维度
	_geoLat string
	// 设备屏幕纵向分辨率，单位：像素
	_screenHeight string
	// 设备屏幕水平分辨率，单位：像素
	_screenWidth string
	// 设备方向：0:未知， 1： 纵向；2： 横向
	_orientation string
	// 运营商，0: 未知， 1:移动,2:电信,3:联通
	_carrier string
	// 网络类型，0：未知，1：WIFI, 2: 2G, 3: 3G, 4: 4G; 5: 5G
	_network string
	// 手机品牌
	_brand string
	// 设备厂商
	_made string
	// 设备型号
	_model string
	// 素材展示所在的页面或者频道
	_posCat string
	// 合作方名称
	_publishName string
	// 给定appid中的位置id，用于提高转化率
	_posId string
	// app类别
	_appCat string
	// app版本号
	_appVer string
	// 对接的appName
	_appName string
	// 对接的appId
	_appId string
	// 关键词的描述
	_description string
	// 关键词
	_keyword string
	// 关键词类型
	_keywordType string
	// 安装转化目标app的标识， 0： 未安装转化目标app, 1: 安装了转化目标app
	_appInstallFlag string
	// 转化目标app, 1：淘宝；2：天猫；3：闲鱼；4：支付宝
	_app string
	// idfa的md5值， 32位小写
	_idfaMd5 string
	// imei的md5值， 32位小写
	_imeiMd5 string
	// idfa原生值
	_idfa string
	// imei原生值
	_imei string
	// 手机号
	_mobile string
	// 转化类型， 1： 激活；2： 新登；32896：定向促活
	_transformType string
	// 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
	_os string
	// 广告创意id，淘宝和媒体协商
	_cid string
	// android最新广告标识
	_oaid string
	// android最新广告标识md5值， 32位小写
	_oaidMd5 string
}

// New
