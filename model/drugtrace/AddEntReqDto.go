package drugtrace

import (
	"sync"
)

// AddEntReqDto 结构体
type AddEntReqDto struct {
	// 企业其他地址省市区信息
	StoreAddrs []Address `json:"store_addrs,omitempty" xml:"store_addrs>address,omitempty"`
	// 企业详细注册地址
	DictRegionDetail string `json:"dict_region_detail,omitempty" xml:"dict_region_detail,omitempty"`
	// 新增企业名称
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 追溯业务负责人联系电话
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 唯一认证代码：唯一认证代码是证件的编号，是国家对您的唯一认证编码。 如：证件是营业执照，唯一认证编码就是营业执照上的社会信用代码或组织机构代码。 证件是身份证，唯一认证编码就是身份证号。 证件是护照，唯一认证编码就是护照编号。 证件是医疗职业许可证，唯一认证编码就是许可证上的卫生机构组织代码或登记号。 以药品注册证为证件入驻的境外企业，唯一认证代码可填“无”。
	OrgCode string `json:"org_code,omitempty" xml:"org_code,omitempty"`
	// 固定电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 追溯业务负责人名称
	ContactPsnNm string `json:"contact_psn_nm,omitempty" xml:"contact_psn_nm,omitempty"`
	// 追溯负责人联系邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 用类型code值，并且需要与userRoleType值相映射：比如userRoleType传的3，只能取枚举值中第三个参数是3的类型；如果没有映射的细分类型，则填0.企业/机构细分类型枚举结构：(类型code，类型名称，对应企业/机构类型code)；枚举值：PRODUCE_CHINA(13, &#34;境内生产企业&#34;, 1), PRODUCE_FOREIGN(14, &#34;境外生产企业&#34;, 1), DISEASE_CENTER(31, &#34;疾控中心&#34;, 3), INOCULATION_ENT(32, &#34;接种单位&#34;, 3), OTHER_MEDICAL_ENT(33, &#34;其他&#34;, 3), MEDICAL_34(34, &#34;综合医院&#34;, 3), MEDICAL_35(35, &#34;中医医院&#34;, 3), MEDICAL_36(36, &#34;中西医结合医院&#34;, 3), MEDICAL_37(37, &#34;民族医院&#34;, 3), MEDICAL_38(38, &#34;专科医院，不含中医专科医院&#34;, 3), MEDICAL_39(39, &#34;疗养院&#34;, 3), MEDICAL_310(310, &#34;护理院(站)&#34;, 3), MEDICAL_311(311, &#34;社区卫生服务中心&#34;, 3), MEDICAL_312(312, &#34;社区卫生服务站&#34;, 3), MEDICAL_313(313, &#34;街道卫生院&#34;, 3), MEDICAL_314(314, &#34;乡镇卫生院&#34;, 3), MEDICAL_315(315, &#34;门诊部&#34;, 3), MEDICAL_316(316, &#34;诊所&#34;, 3), MEDICAL_317(317, &#34;卫生所(室)&#34;, 3), MEDICAL_318(318, &#34;医务室&#34;, 3), MEDICAL_319(319, &#34;中小学卫生保健所&#34;, 3), MEDICAL_320(320, &#34;村卫生室&#34;, 3), MEDICAL_321(321, &#34;急救中心&#34;, 3), MEDICAL_322(322, &#34;急救中心站&#34;, 3), MEDICAL_323(323, &#34;急救站&#34;, 3), MEDICAL_324(324, &#34;血站&#34;, 3), MEDICAL_325(325, &#34;单采血浆站&#34;, 3), MEDICAL_326(326, &#34;妇幼保健院&#34;, 3), MEDICAL_327(327, &#34;妇幼保健所包括妇女、儿童保健所&#34;, 3), MEDICAL_328(328, &#34;妇幼保健站包括妇幼保健中心&#34;, 3), MEDICAL_329(329, &#34;生殖保健中心&#34;, 3), MEDICAL_330(330, &#34;专科疾病防治院&#34;, 3), MEDICAL_331(331, &#34;专科疾病防治所(站、中心)&#34;, 3), MEDICAL_333(333, &#34;卫生防疫站&#34;, 3), MEDICAL_334(334, &#34;卫生防病中心&#34;, 3), MEDICAL_335(335, &#34;预防保健中心&#34;, 3), MEDICAL_336(336, &#34;卫生监督所(局)&#34;, 3), MEDICAL_337(337, &#34;卫生(综合)监督检验(监测、检测)所(站)&#34;, 3), MEDICAL_338(338, &#34;环境卫生监督检验(监测、检测)所(站)&#34;, 3), MEDICAL_339(339, &#34;放射卫生监督检验(监测、检测)所(站)&#34;, 3), MEDICAL_340(340, &#34;劳动(职业、工业)卫生监督检验(监测、检测)所(站)&#34;, 3), MEDICAL_341(341, &#34;食品卫生监督检验(监测、检测)所(站)&#34;, 3), MEDICAL_342(342, &#34;学校卫生监督检验(监测、检测)所(站)&#34;, 3), MEDICAL_343(343, &#34;其他卫生监督检验(监测、检测)所(站)&#34;, 3), MEDICAL_344(344, &#34;医学科学(研究)院（所）&#34;, 3), MEDICAL_345(345, &#34;预防医学研究院（所）&#34;, 3), MEDICAL_346(346, &#34;中医(药)研究院(所)&#34;, 3), MEDICAL_347(347, &#34;中西医结合研究所&#34;, 3), MEDICAL_348(348, &#34;民族医(药)学研究所&#34;, 3), MEDICAL_349(349, &#34;医学专科研究所&#34;, 3), MEDICAL_350(350, &#34;药学研究所（包括药用植物研究所）&#34;, 3), MEDICAL_351(351, &#34;医学普通高中等学校&#34;, 3), MEDICAL_352(352, &#34;医学成人学校&#34;, 3), MEDICAL_353(353, &#34;医学在职培训机构包括各类卫生技术人员、管理人员培训中心等&#34;, 3), MEDICAL_354(354, &#34;健康教育所包括健康教育研究所&#34;, 3), MEDICAL_355(355, &#34;健康教育站(中心)包括卫生宣教馆&#34;, 3), MEDICAL_356(356, &#34;临床检验中心（所、站）&#34;, 3), MEDICAL_357(357, &#34;卫生新闻出版社&#34;, 3), MEDICAL_358(358, &#34;其他卫生事业机构&#34;, 3), MEDICAL_359(359, &#34;红十字会&#34;, 3), MEDICAL_360(360, &#34;医学会包括预防医学会、护理学会、医院管理学会等各类卫生专业学会&#34;, 3), MEDICAL_361(361, &#34;协会&#34;, 3), MEDICAL_362(362, &#34;其他卫生社会团体&#34;, 3), CHAIN_DRUG_STORE(41, &#34;连锁药店总部&#34;, 4), DIRECT_CHAIN_DRUGS_STORE(42, &#34;直营连锁药店&#34;, 4), JOIN_CHAIN_DRUG_STORE(43, &#34;加盟连锁药店&#34;, 4), SINGLE_DRUG_STORE(44, &#34;单体药店&#34;, 4), OTHER_DRUG_STORE(45, &#34;其他&#34;, 4), MARKET_AUTH_HOLD_CHAINA(71, &#34;境内持证商&#34;, 7), MARKET_AUTH_HOLD_FOREIGN(72, &#34;境外持证商&#34;, 7) ;
	EntOrgType int64 `json:"ent_org_type,omitempty" xml:"ent_org_type,omitempty"`
	// 用类型code值：比如是医疗机构则传3.企业/机构类型枚举结构：(类型code，类型名称)； 枚举值：PROD_ENT(1,&#34;生产企业&#34;), WHOLESALE_ENT(2,&#34;批发企业&#34;), MEDICAL_ENT(3,&#34;医疗机构&#34;), RESALE_ENT(4,&#34;零售企业&#34;), LOGISTICS_ENT(5,&#34;物流企业&#34;), MAH_ENT(7,&#34;上市许可持有人&#34;), CENTRAL_RANDOMIZATION_SYSTEM_PROVIDER(8,&#34;中央随机化系统提供商&#34;);
	UserRoleType int64 `json:"user_role_type,omitempty" xml:"user_role_type,omitempty"`
	// 企业注册地址省市区信息
	RegAddr *Address `json:"reg_addr,omitempty" xml:"reg_addr,omitempty"`
	// 用类型code值：比如营业执照资质则传8.资质证照类型枚举结构：(类型code，类型值)；枚举值：DRUG_BUSINESS_LICENSE(1, &#34;药品经营许可证&#34;), GSP(2, &#34;GSP&#34;), PRACTICE_LICENSE_OF_MEDICAL_INSTITUTION(3, &#34;医疗机构执业许可证&#34;), GMP(4, &#34;GMP&#34;), PHARMACEUTICAL_PRODUCTION_LICENSE(5, &#34;药品生产许可证&#34;), GUP(6, &#34;GUP&#34;), REGISTRATION_CERTIFICATE_OF_FOREIGN_ENTERPRISES(7, &#34;外国企业常驻中国代表机构登记证&#34;), BUSINESS_LICENSE(8, &#34;营业执照&#34;), OTHER(9, &#34;其他&#34;), ORG_CODE_LICENSE(10, &#34;组织机构代码证&#34;), PUBLIC_INSTITUTIONS_LICENSE(11, &#34;事业单位法人证书&#34;), PRIVATE_NON_ENTERPRISE_ORGANIZATION_LICENSE(12, &#34;民办非企业单位登记证书&#34;), AGENCY_BUSINESS_LICENSE(13, &#34;代理公司营业执照&#34;), IMPORT_DRUG_REGISTRATION_LICENSE(14, &#34;进口药品注册证&#34;), DRUG_REGISTRATION_APPROVAL_LICENSE(15, &#34;药品注册批件&#34;), AUTHORIZATION_LICENSE(16, &#34;授权书&#34;) ;
	LicenseType int64 `json:"license_type,omitempty" xml:"license_type,omitempty"`
}

var poolAddEntReqDto = sync.Pool{
	New: func() any {
		return new(AddEntReqDto)
	},
}

// GetAddEntReqDto() 从对象池中获取AddEntReqDto
func GetAddEntReqDto() *AddEntReqDto {
	return poolAddEntReqDto.Get().(*AddEntReqDto)
}

// ReleaseAddEntReqDto 释放AddEntReqDto
func ReleaseAddEntReqDto(v *AddEntReqDto) {
	v.StoreAddrs = v.StoreAddrs[:0]
	v.DictRegionDetail = ""
	v.EntName = ""
	v.MobilePhone = ""
	v.OrgCode = ""
	v.Tel = ""
	v.ContactPsnNm = ""
	v.Email = ""
	v.EntOrgType = 0
	v.UserRoleType = 0
	v.RegAddr = nil
	v.LicenseType = 0
	poolAddEntReqDto.Put(v)
}
